package main

import "fmt"

func main() {
	a := 5
	if a < 10 && a > 2 {
		fmt.Println("a는 10작고 2보다 크다")
	} else {
		fmt.Println("a는 10보다 크거나, a는 2보다 작다")
	}
}
