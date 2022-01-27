package structs
import "fmt"

//map보다 유연하고 object같음
type person struct {
  name string
  age int
  favFood []string
}
func Structs(){
  favFood:= []string{"초밥", "라면"}
  leesky := person{"leesky", 23, favFood}
  
  leesky2 := person{
    name: "leehaneul",
    age: 23,
    favFood: favFood,
  }

  fmt.Println(leesky ,leesky.favFood, leesky2)  
}