//1. main 패키지
//컴파일러는 main 패키지부터 컴파일 한다.
//function main 을 가장 먼저 실행한다.
//fmt => 포맷팅하는 것으로 추측
//go run main.go => 실행

//2. Packages and imports
//export function은 대문자로 시작함
//소문자로 시작하면 private임

//3. 변수와 상수
// 상수 : const name type = "leesky"
// 변수 : var name type = ""
//   -> name:="leesky" <= 동적 타입언어처럼 타입을 지정해준다.
//   -> 위와 같은 축약 선언은 함수내부에서만 작동가능
//   -> 변수에만 적용가능하다
package main

import (
	"fmt"
	"main/something"
)

func main() {
	fmt.Println("Hello, World!")
  something.SayHello()
  const name string = "leesky";


}