package com.example.ssechat.controller;

import com.example.ssechat.service.ChatService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

import java.io.IOException;

@CrossOrigin(origins = "*")  // CORS 설정 추가
@RestController
@Slf4j
@RequestMapping("/chat")
public class ChatController {

    private final ChatService chatService;

    public ChatController(ChatService chatService) {
        this.chatService = chatService;
    }

    @GetMapping("/subscribe/{username}")
    public SseEmitter subscribe(@PathVariable String username) {
        return chatService.subscribe(username);
    }

    @PostMapping("/send")
    public void sendMessage(@RequestParam String username, @RequestParam String message) {
        chatService.sendMessage(username, message);
    }

    @ExceptionHandler(IOException.class)
    public void handleException(IOException e) {
        log.warn("IOException occurred: {}", e.getMessage());
    }
}
