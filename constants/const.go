package main

import "fmt"

/*
상수는 const 키워드와 함께 변수처럼 선언합니다.
상수는 문자(character), 문자열(string), 부울(boolean), 숫자 타입 중의 하나가 될 수 있습니다.
*/
const Pi = 3.14

func main() {
	const World = "안녕"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
