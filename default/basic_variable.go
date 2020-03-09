package main

import (
	"fmt"
	"math/cmplx"
)

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8의 다른 이름(alias)

rune // int32의 다른 이름(alias)
     // 유니코드 코드 포인트 값을 표현합니다.

float32 float64

complex64 complex128
*/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//fmt의 인자를 변수의 type, value값으로 포멧팅하여 출력
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
