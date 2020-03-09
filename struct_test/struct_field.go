package main

import "fmt"

/*
구조체 필드
	구조체에 속한 필드(데이터)는 dot(.) 으로 접근합니다.
*/
type Vertex struct {
	X int
	Y int
}

func main() {
	//struct에 값 할당
	v := Vertex{1, 2}

	//struct 변수값 변경
	v.X = 4

	//struct 변수값 출력
	fmt.Println(v.X)
	//4
}
