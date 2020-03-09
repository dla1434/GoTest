package main

import "fmt"

/*
구조체 (Structs)
	struct는 필드(데이터)들의 조합입니다.
	(그리고 type 선언으로 struct의 이름을 지정할 수 있습니다.)
*/
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	//{1 2}
}
