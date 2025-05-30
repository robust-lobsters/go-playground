package basic

import "fmt"

// = public 접근자로 선언된 함수
func PublicFunc() {
	fmt.Println("PublicFunc 호출")

	// 변수 선언
	var x = 1
	y := 2
	// 1 2
	fmt.Println(x, y)

	// 상수 선언
	const constant = 41

	// 배열 선언
	arr := [3]int{1, 2, 3}
	// 41 2
	fmt.Println(constant, arr[1])

	// <key: string / value: int> 맵 선언
	// {"a": 1, "b": 2}로 초기화
	mapExample := map[string]int{"a": 1, "b": 2}

	// 0 -> 키가 없을 시 자료형의 기본값
	fmt.Println(mapExample["A"])
	// 2
	fmt.Println(mapExample["b"])

	// User 구조체 선언
	type User struct {
		Name string
		Age  int
	}

	user := User{Name: "이름", Age: 20}

	// {이름 20}
	fmt.Println(user)
}
