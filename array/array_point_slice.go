package main

import "fmt"

/*
슬라이스 자르기 (Slicing slices)
	슬라이스는 재분할 할 수도 있고, 같은 배열을 가리키는(point) 새로운 슬라이스를 만들 수 도 있습니다.
	예제로 살펴보면
	s[lo:hi]

	위의 표현은 lo 에서 hi-1 의 요소(element)를 포함하는 슬라이스입니다. 따라서
	s[lo:lo]
	는 빈(empty) 슬라이스 이고
	s[lo:lo+1]
	는 하나의 요소를 가집니다.
*/
func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	//p == [2 3 5 7 11 13]
	fmt.Println("p[1:4] ==", p[1:4])
	//p[1:4] == [3 5 7]

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])
	//p[:3] == [2 3 5]

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])
	//p[4:] == [11 13]
}
