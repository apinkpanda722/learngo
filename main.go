package main

import (
	"fmt"
	"learngo/mydict"
)

// accounts
//func main() {
//	account := accounts.NewAccount("nico")
//	account.Deposit(10)
//	err := account.Withdraw(5)
//	fmt.Println(account)
//	account.ChangeOwner("tale")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(account.Balance(), account.Owner())
//	}
//}

func main() {
	//dictionary := mydict.Dictionary{"first": "First word"}
	//definition, err := dictionary.Search("first")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(definition)
	//}

	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search("hello")
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
