# Server Sent Event 실습
- 서버에서 클라이언트로 실시간 이벤트를 전달하는 웹 기술
- SSE는 단방향 통신이며 클라이언트의 별도 추가요청 없이 서버에서 업데이트를 스트리밍할 수 있다는 특징

## 장점
1. HTTP를 통해 통신하므로 다른 프로토콜은 필요가 없고, 구현이 굉장히 쉽다는 것이다.
2. 네트워크 연결이 끊겼을 때 자동으로 재연결을 시도한다.
3. 실시간으로 서버에서 클라이언트로 데이터를 전송할 수 있다. 폴링 같은 경우는 실시간이라고 보기 어려운 점이 있는데, 이러한 한계를 극복한다.

## 단점
1. GET 메소드만 지원하고, 파라미터를 보내는데 한계가 있다.
2. 단방향 통신이며, 한 번 보내면 취소가 불가능하다는 단점이 있다.
3. 클라이언트가 페이지를 닫아도 서버에서 감지하기가 어렵다는것도 단점이다.
4. SSE는 지속적인 연결을 유지해야 하므로, 많은 클라이언트가 동시에 연결을 유지할 경우 서버 부담이 커질 수 있다.

## 구현

### 구독
```java
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
```
- 타임아웃
  - 클라이언트 측에서 일정 시간 동안 서버로부터 데이터를 받지 못할 경우에 발생하는 상황을 말한다. 
  - 타임아웃이 발생하면 브라우저에서 자동으로 서버에 재연결 요청을 보내서 해결하게 된다

### 전송
```java
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
```
