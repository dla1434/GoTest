package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

/*
레인지 (Range, 범위)
	for 반복문에서 range 를 사용하면 슬라이스나 맵을 순회(iterates)할 수 있습니다.
*/
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
		/*
			2**0 = 1
			2**1 = 2
			2**2 = 4
			2**3 = 8
			2**4 = 16
			2**5 = 32
			2**6 = 64
			2**7 = 128
		*/
	}
}
