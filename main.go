//1. main 패키지
//컴파일러는 main 패키지부터 컴파일 한다.
//function main 을 가장 먼저 실행한다.
//fmt => 포맷팅하는 것으로 추측
//go run main.go => 실행

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}