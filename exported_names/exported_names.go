package main

import (
	"fmt"
	"math"
)

/*
패키지를 Import 하면 패키지가 외부로 export 한 것들(메서드나 변수, 상수 등)에 접근할 수 있습니다.
Go에서는 첫 문자가 대문자로 시작하면 그 패키지를 사용하는 곳에서 접근할 수 있는 exported name이 됩니다.
예를 들어 Foo 와 FOO 는 외부에서 참조할 수 있지만 foo 는 참조 할 수 없습니다.
예제를 실행해보세요. 에러가 발생한다면 math.pi 를 math.Pi 로 변경 해보세요.
*/
func main() {
	//fmt.Println(math.pi) : 에러 발생
	fmt.Println(math.Pi)
}
