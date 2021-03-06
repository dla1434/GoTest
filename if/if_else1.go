package main

import (
	"fmt"
	"math"
)

//if 에서 짧은 명령문을 통해 선언된 변수는 else 블럭 안에서도 사용할 수 있습니다.
// 주의점) if 다음에 else를 } 뒤에 쎃야 정상 인식된다...} 다음라인에 else를 작성하면 에러가 발생되니 주의
func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		//27 >= 20
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	//9 20
}
