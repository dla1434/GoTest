package main

import "fmt"

/*
   맵 리터럴 (Map literals)
   	맵 리터럴은 구조체 리터럴과 비슷하지만 key 를 반드시 지정해야 합니다.
*/
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
	//map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}
