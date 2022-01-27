package pointer

import "fmt"

func Pointer(){
  // 값에 의한 복사
  a :=2;

  b :=&a;
  a = 10;
  *b = 1



  //주소 연산자
  fmt.Println(&a, b, a, *b)
}
