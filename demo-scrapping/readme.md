# Go
- 라이브러리 없는 거 찾아서 다운
```text
go mod tidy
```

- 명시적인 타입 선언: 변수의 이름 뒤에 타입을 명시합니다.
```go
var 변수명 타입
```
- 타입 추론: 변수명 := 값

- 예시
```go
var age int           // int 타입의 age 변수를 선언
var name string       // string 타입의 name 변수를 선언
var temperature float64   // float64 타입의 temperature 변수를 선언

age = 30             // age 변수에 값 할당
name = "John"        // name 변수에 값 할당
temperature = 25.5   // temperature 변수에 값 할당

// 또는
height := 180        // int 타입의 height 변수를 선언하고 값 할당 (타입 추론)
```
