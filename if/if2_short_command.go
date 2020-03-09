package main

import (
	"fmt"
	"math"
)

/*
If와 짧은 명령 사용하기
	for 처럼 if 에서도 조건문 앞에 짧은 문장을 실행할 수 있습니다.
	예제에서는 조건을 비교하기 전에 v := math.Pow(x,n) 을 실행했습니다.
	(예제 코드의 pow 함수에서 return 전에 v 를 사용해보세요.)
	짧은 실행문을 통해 선언된 변수는 if 안쪽 범위(scope) 에서 만 사용할 수 있습니다.
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
