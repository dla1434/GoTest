package main

import "fmt"

func main() {
	//C와 Java에서 처럼 전.후 처리를 제외하고 조건문만 표현할 수도 있습니다.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
