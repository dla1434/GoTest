package main

import "fmt"

/*
맵 다루기 (Mutating Maps)
	1. 맵 m 의 요소를 삽입하거나 수정하기:
		m[key] = elem

	2. 요소 값 가져오기:
		elem = m[key]

	3. 요소 지우기:
		delete(m, key)

	4. 키의 존재 여부 확인하기:
		elem, ok = m[key]

		> 위의 ok 의 값은 m 에 key 가 존재한다면 true 존재하지 않으면 false , elem 은 타입에 따라 0(zero value) 가 됩니다.
		> 이처럼 map 을 읽을 때, 존재하지 않는 key 의 반환 값은 타입에 맞는 zero value 입니다.
*/
func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	//The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	//The value: 48

	v1, ok := m["Answer"]
	fmt.Println("The value:", v1, "Present?", ok)
	//The value: 48 Present? true
	//값이 있다면 v1에 값이 있고, ok변수에는 true 값이 들어간다.

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	//The value: 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	//The value: 0 Present? false
	//값이 없다면 v에 값이 0이고, ok변수에는 false 값이 들어간다.
}
