package main

import "fmt"

/*
구조체 리터럴 (Struct Literals)
	구조체 리터럴은 필드와 값을 나열해서 구조체를 새로 할당하는 방법입니다.
	원하는 필드를 `{Name: value}` 구문을 통해 할당할 수 있습니다. (필드의 순서는 상관 없습니다.)
	특별한 접두어 & 를 사용하면 구조체 리터럴에 대한 포인터를 생성할 수 있습니다.
*/
type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
	//{1 2} &{1 2} {1 0} {0 0}
}
