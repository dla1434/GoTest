package main

import "fmt"

func main() {
	sum := 0
	//기본적인 for 반복문은 C와 Java 언어와 거의 유사합니다. 다른점은 소괄호 ( )가 필요하지 않다는 것입니다.
	//하지만 실행문을 위한 중괄호 { } 는 필요합니다.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
