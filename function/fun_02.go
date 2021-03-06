package main

import "fmt"

/*
두 개 이상의 매개변수가 같은 타입(type)일 때, 같은 타입을 취하는 마지막 매개변수에만 타입을 명시하고 나머지는 생략할 수 있습니다.
예를 들어
x int, y int
과 같은 문장은 아래와 같이 표현할 수 있습니다.
x, y int
*/
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
