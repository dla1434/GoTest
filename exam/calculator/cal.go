package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("첫번째 숫자를 입력하세요")

	//콘솔 입력값을 받아오기 위해 bufio를 사용한다.
	reader := bufio.NewReader(os.Stdin)

	//한 줄을 읽어오는데..\n 즉 개행문자가 오기 전까지 읽어들려라.
	//_는 이름 없는 변수인데..에러가 들어오는데..처리하지 않을거라서..명시적으로 _를 사용
	line, _ := reader.ReadString('\n')

	//line 변수를 trimSpace로 앞 뒤의 공백 제거
	line = strings.TrimSpace(line)
	n1, _ := strconv.Atoi(line)

	fmt.Println("두번째 숫자를 입력하세요")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2, _ := strconv.Atoi(line)

	fmt.Println("입력하신 숫자는 %d, %d 입니다.", n1, n2)
	fmt.Println("연산자를 입력하세요")
	line, _ = reader.ReadString('\n')
	op := strings.TrimSpace(line)

	if op == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if op == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if op == "*" {
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	} else if op == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}
}
