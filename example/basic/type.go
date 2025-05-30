package basic

import "fmt"

// = public 접근자로 선언된 함수
func PublicFunc() {
	fmt.Println("PublicFunc 호출")

	// 변수 선언
	var x = 1
	y := 2

	fmt.Println(x, y)

	// 상수 선언
	const constant = 41

	// 배열 선언
	arr := [3]int{1, 2, 3}

	fmt.Println(constant, arr[1])
}
