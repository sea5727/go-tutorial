# context-tutorial

```go
// 컨텍스트 생성
ctx := context.Background()

// 컨텍스트에 값 추가
// context.WithValue 함수를 사용하여 새로운 컨텍스트를 생성함
ctx = context.WithValue(ctx, "current_user", currentUser)

// 함수 호출시 컨텍스트를 파라미터로 전달
myFunc(ctx)

```