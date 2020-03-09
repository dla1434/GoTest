package main

import "fmt"

/*
new 함수
	new(T) 는 모든 필드가 0(zero value) 이 할당된 T 타입의 포인터를 반환합니다.
	( zero value 는 숫자 타입에서는 0 , 참조 타입에서는 nil 을 뜻합니다 )

	var t *T = new(T)
	또는
	t := new(T)

	위의 변수 t는 T 에서 반환된 포인터를 가집니다.
*/
type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	//&{0 0}

	v.X, v.Y = 11, 9
	fmt.Println(v)
	//&{11 9}
}
