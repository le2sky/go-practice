package main

import (
	"fmt"
	"main/mydict"
)

func main(){
  dictionary := mydict.Dictionary{}
  baseWord := "hello"
  dictionary.Add(baseWord, "First")
  err := dictionary.Update(baseWord, "Second")
  if err != nil {
    fmt.Println(err)
  }

  word, _ := dictionary.Search(baseWord)
  fmt.Println(word)
}