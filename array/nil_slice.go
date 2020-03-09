package main

import "fmt"

/*
빈 슬라이스
	슬라이스의 zero value는 nil 입니다.
	nil 슬라이스는 길이와 최대 크기가 0입니다.
	(슬라이스에 대해 더 알고 싶다면 다음 글을 읽어보세요. Slices: usage and internals)
*/
func main() {
	var z []int

	fmt.Println(z, len(z), cap(z))
	//[] 0 0

	if z == nil {
		fmt.Println("nil!")
		//nil!
	}
}
