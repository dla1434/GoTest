package main

import "fmt"

/*
함수는 매개변수(인자)를 가질 수 있습니다.
예를 들어 add 라는 함수는 두개의 int 타입 매개변수를 받습니다.
C, C++, Java 언어와 다르게 매개변수의 타입은 변수명 뒤에 명시합니다.
(타입을 왜 변수명 뒤에 명시하는지에 대한 자세한 내용은 Go's declaration syntax를 참고하시기 바랍니다.)
(간단히 설명하면 코드를 왼쪽에서 오른쪽으로 읽을 때 자연스럽게 읽기 위해서 입니다.)
*/
//() 뒤에 int는 리턴 타입이다.
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
