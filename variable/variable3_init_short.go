package main

import "fmt"

/*
변수의 짧은 선언
	변수 선언 시 var와 변수 타입(e.g. int, bool)을 생략할 수 있다.
	변수 선언 시 := 사용하여 값을 넣으면 변수값에 따라 변수타입이 지정된다.
	ex) x := 3는 var x int = 3과 동일하다.
	(주의점은 함수 밖에서는 := 선언을 사용할 수 없습니다.)
*/
func main() {
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)

	d := 3
	//아래처럼 사용도 가능
	//var d = 3
	//var d int = 3
	fmt.Println("d is : ", d)
}
