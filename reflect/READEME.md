# Reflect


[golang doc문서](https://golang.org/pkg/reflect/)   
   
[참조 블로그](https://johngrib.github.io/wiki/golang-reflect/)   




reflect는 run-time동안 프로그램이 임의의 타입을 사용하는데 쓰인데. 일반적으로 interface{}의 Value를 사용하여, TypeOf을 호출한다.

ValueOf 은 run-time의 데이터 표현식을 리턴한다. 

```go
package main

import (
	"fmt"
	"reflect"
)

func typeSwitch(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is int.")
	case string:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is string.")
	default:
		fmt.Println("type:", reflect.ValueOf(value).Type())
	}
}

func main() {
	fmt.Println("reflect.go")
	typeSwitch(42)
	typeSwitch("42")
	typeSwitch([]float32{3.14})
}
```

```
type: int
42 is int.
type: string
42 is string.
type: []float32
```