package main

import "fmt"

//함수는 매개변수(인자)를 가질 수 있습니다.
//예를 들어 add 라는 함수는 두개의 int 타입 매개변수를 받습니다.
//C, C++, Java 언어와 다르게 매개변수의 타입은 변수명 뒤에 명시합니다.
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
