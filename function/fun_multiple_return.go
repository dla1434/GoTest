package main

import "fmt"

//하나의 함수는 여러 개의 결과를 반환할 수 있습니다.
//string의 s는 대문자가 아니 소문자이다.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "World")
	fmt.Println(a, b)
}
