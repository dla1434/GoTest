package main

import "fmt"

func main() {
	//이전의 예제에서 처럼 조건문만 표시하면 C언어에서 while 을 사용하듯 for 를 사용할 수 있습니다.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
