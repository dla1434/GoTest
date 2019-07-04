package main

import "fmt"

func main() {
	//변수의 짧은 선언
	//함수 내에서 := 을 사용하면 var 과 명시적인 타입(e.g. int, bool) 을 생략할 수 있습니다.
	//(그러나 함수 밖에서는 := 선언을 사용할 수 없습니다.)
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}
