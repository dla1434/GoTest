package main

import "fmt"

/*
숫자형 상수 (Numeric Constants)
	숫자형 상수는 정밀한 값(values) 을 표현할 수 있습니다.
	타입을 지정하지 않은 상수는 문맥(context)에 따라 타입을 가지게 됩니다.
	코드를 통해 needInt(Big) 는 어떤 결과를 출력할지 실험해보세요.
*/
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	//21
	fmt.Println(needFloat(Small))
	//0.2
	fmt.Println(needFloat(Big))
	//1.2676506002282295e+29
}
