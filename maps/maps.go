package maps

import "fmt"

func Maps(){
  leesky := map[string]string{"name":"leesky", "age": "12"}
  fmt.Println(leesky)
  for _, value := range leesky {
    fmt.Println(value)
  }

}