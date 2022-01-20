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
//   -> 변수에만 적용가능하다 (주의: 선언이라는 개념임)

//4. function
// 아래 주석 확인


//5. loop
// go는 for만 사용 가능


//6.if with a twist


package main

import (
	"fmt"
	"main/something"
	"strings"
)

func canIDring(age int) bool{
  //variable expression
  if koreanAge := age + 2; koreanAge < 18 {
    return false
  } 
  return true

}


func superAdd(numbers ...int) (result int) {
  for _,v := range numbers {
    result+=v
  }
  return
}

// Naked return + defer 특이함;
func lenAndUpperNaked(name string) (length int, uppercase string){
  defer fmt.Println("done")
  length = len(name)
  uppercase = strings.ToUpper(name);
  return 
}

// 뒤에 꺼만 하면 앞에는 뒤에 따라감!
func mult(a ,b int) int{
  return a * b
}

//다중 리턴 특이하다.
func lenAndUpper(name string) (int, string){
  return len(name), strings.ToUpper(name)
}

//js rest이랑 비슷함 타입은 지정한 놈의 배열
func repeatMe(words ...string){
  fmt.Println(words)
}

func main() {
	fmt.Println("Hello, World!")
  fmt.Println(mult(3, 3))

  totalLength, upperName := lenAndUpper("leesky")
  fmt.Println(totalLength, upperName)

  repeatMe("lee", "sky", "study", "go")

  fmt.Println(canIDring(16));

  total := superAdd(1,2,3,4,5)
  fmt.Println(total)

  something.SayHello()
  const name string = "leesky"
}