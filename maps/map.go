package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

/*
맵 (Maps)
	맵은 값에 키를 지정합니다.
	맵은 반드시 사용하기 전에 make 를 명시해야합니다. (주의: new 가 아닙니다)
	make 를 수행하지 않은 nil 에는 값을 할당할 수 없습니다.
*/
func main() {
	/*
		자바의 hashmap을 사용하는데..
		map[string] : 키값으로는 string이 들어가고
		map[string]Vertex : 값으로는 struct가 들어가는 형태이이다.
		또 make를 사용하여 생성해야 한다.
	*/
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	//{40.68433 -74.39967}

	fmt.Println(m["Bell Labs"].Lat)
	//40.68433
}
