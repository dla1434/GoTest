package main

import (
	"fmt"
	"math"
)

func main() {
	//%v - 인자타입에 맞게 값을 출력
	fmt.Printf("Test %v \n", 3)

	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}
