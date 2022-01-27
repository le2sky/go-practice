package array

import "fmt"

func Array(){
  names := [5]string{"lee", "haneul"}
  //slice는 기본적으로 array 지만 길이 제한이 없음
  namesWithSlice := []string{"lee", "haneul"}
  newSlice := append(namesWithSlice, "0623");
  


  names[2] = "go";
  names[3] = "lang";
  names[4] = "study";
  
  fmt.Println(names,namesWithSlice, newSlice)  
}

