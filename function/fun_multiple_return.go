package main

import "fmt"

/*
하나의 함수는 여러 개의 리턴 결과를 반환할 수 있습니다.
예제 코드에서의 함수는 두개의 문자열을 반환합니다.
string의 s는 대문자가 아니 소문자이다.
*/
//swap을 사용하여 출력 순서를 바뀌는 함수
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "World")
	fmt.Println(a, b)
	//world hello
}
