package main

import (
	"example/basic"
	"fmt"
)

func main() {
	privateFunc()
	basic.PublicFunc()
}

// = private 접근자로 선언된 함수
func privateFunc() {
	fmt.Println("privateFunc 호출")
}
