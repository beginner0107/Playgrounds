package com.example.ssechat.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

import java.io.IOException;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.Collectors;

@Slf4j
@Service
public class ChatService {
    private final Map<String, SseEmitter> emitters = new ConcurrentHashMap<>();

    public SseEmitter subscribe(String username) {
        SseEmitter emitter = new SseEmitter(30 * 60 * 1000L); // 30분 타임아웃 설정
        emitters.put(username, emitter);
        emitter.onCompletion(() -> {
            emitters.remove(username);
            log.info("Emitter completed for user: {}", username);
        });
        emitter.onTimeout(() -> {
            emitters.remove(username);
            log.info("Emitter timed out for user: {}", username);
        });
        emitter.onError((e) -> {
            emitters.remove(username);
            log.info("Emitter error for user: {}", username, e);
        });
        return emitter;
    }

    public void sendMessage(String username, String message) {
        log.info("Sending message from user {}: {}", username, message);
        List<String> toRemove = emitters.entrySet().stream().filter(entry -> {
            try {
                synchronized (entry.getValue()) {
                    entry.getValue().send(SseEmitter.event().name("chat").data(username + ": " + message));
                    log.info("Message sent to user {}: {}", entry.getKey(), message);
                }
                return false;
            } catch (IOException | IllegalStateException e) {
                log.error("Failed to send message to {}: {}", entry.getKey(), e.getMessage(), e);
                return true;
            }
        }).map(Map.Entry::getKey).collect(Collectors.toList());

        toRemove.forEach(emitters::remove);
    }
}
