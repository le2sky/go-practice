package main

import (
	"fmt"
	"main/accounts"
)

func main(){
  account := accounts.NewAccount("leesky")
  fmt.Println(account)
}