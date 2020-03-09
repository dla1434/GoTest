package main

import "fmt"

/*
레인지 (2)
	_ 를 이용해서 인덱스(index)나 값(value)를 무시할 수 있습니다.
	만약 인덱스만 필요하다면 `, value` 부분을 다 제거하면 됩니다.

	for i, value := range pow {
		pow[i] = 1 << uint(i)
	}
	에서

	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	처럼 사용할 수 있습니다.
*/
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
		/*
			1
			2
			4
			8
			16
			32
			64
			128
			256
			512
		*/
	}
}
