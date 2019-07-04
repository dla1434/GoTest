package main

import "fmt"

//변수 선언과 함께 변수 각각을 초기화를 할 수 있습니다.
var x, y, z int = 1, 2, 3

//초기화를 하는 경우 타입(type)을 생략할 수 있습니다. 변수는 초기화 하고자 하는 값에 따라 타입이 결정됩니다.
//var c,python,java bool = true, false, "no!"
//bool type을 제외하면 변수값을 보고 type이 알아서 지정된다.
var c, python, java = true, false, "no!"

func main() {
	fmt.Println(x, y, z, c, python, java)
}
