package main

import (
	"fmt"
	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	//account.Deposit(10)
	//err := account.Withdraw(5)
	//account.ChangeOwner("tale")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(account.Balance(), account.Owner())
	//}
}
