package main

import "fmt"

func main() {
	a := 4
	if a == 3 {
		fmt.Println("a는 3")
	} else if a == 4 {
		fmt.Println("a는 4")
	} else {
		fmt.Println("a는 3, 4도 아니다.")
	}
}
