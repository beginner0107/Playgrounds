# localhost에서 https 이용하기
## 사전 지식
- 자바에서는 다음 두 가지 인증서 형식을 지원
  - PKCS12(Public Key Crytographic Standard #12) : 여러 인증서와 키를 포함할 수 있으며, 암호로 보호된 형식(많이 사용)
  - JKS(Java KeyStore) : PKCS12와 유사. 독점 형식이며, Java 환경으로 제한
- https://github.com/FiloSottile/mkcert

## mkcert 문서를 참고하여 설치
- window
```text
choco install mkcert
```
## https를 적용할 프로젝트 최상위 디렉토리에서 다음 명령어 실행
```text
mkcert -install
mkcert -pkcs12 localhost 127.0.0.1 # localhost.p12 파일 생성됨.

mkcert localhost # `#*-key.pem`, `*.pem` 파일이 생성됨
```
```yml
server:
  port: 8080
  ssl:
    enabled: true
    key-store: classpath:cert/www.zoos.dev.p12
    key-store-type: PKCS12
    key-store-password: changeit
  http2:
    enabled: true #openfeign 오류 시
```

## 적용 확인
- 브라우저에서 localhost로 접속한 후 좌물쇠 모양 아이콘이 나타나는지 확인

## hosts 파일 변경해서 도메인처럼 사용하기
- 윈도우 기준
- C:\Windows\System32\drivers\etc\host파일
- 127.0.0.1 -> 
```text
127.0.0.1 localhost - 기존
127.0.0.1 www.zoos.dev - 변경
```
- 다시한번 더 mkcert
```text
mkcert -pkcs12 www.zoos.dev
```
