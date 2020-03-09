package main

import "fmt"

/*
이름이 정해진 결과 (Named results)
	함수는 매개변수를 취합니다.
	Go에서 함수는 여러 개의 결과를 반환할 수 있습니다.
	반환 값에 이름을 부여하면 변수처럼 사용할 수도 있습니다.
	결과에 이름을 붙히면, 반환 값을 지정하지 않은 return 문장으로 결과의 현재 값을 알아서 반환합니다.
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	//7 10
}
