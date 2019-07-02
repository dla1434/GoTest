package main

import(
	"fmt"
	"math"
)

func main(){
	fmt.Println("Hello, 안녕", math.Pi, "Day")

	fmt.Printf("Now you have %g problems.", math.Nextafter(2,3))
}