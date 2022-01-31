package main

import (
	"fmt"
	"main/accounts"
)

func main(){
  account := accounts.NewAccount("leesky")

  account.Deposit(10)
  err := account.Withdraw(20)
  if err != nil {
    fmt.Println(err)
    //log.Fatalln(err) kill the program
  }
  fmt.Println(account.Balance(), account.Owner())
  fmt.Println(account)
}