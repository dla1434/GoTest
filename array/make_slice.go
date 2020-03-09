package main

import "fmt"

/*
슬라이스 만들기 :: 배열 크기 지정이라고 보면 된다.
	슬라이스는 make 함수로 만들 수 있습니다.
	이렇게 생성된 슬라이스는 0을 할당한 배열을 생성하고, 그것을 참조(refer)합니다.

	a := make([]int, 5)  // len(a)=5

	make 함수의 세번째 매개변수로 용량(capacity)를 제한할 수 있습니다.
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
*/
func main() {
	a := make([]int, 5)
	printSlice("a", a)
	//a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b)
	//b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c)
	//c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d)
	//d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
