package main

import (
	"fmt"

	"github.com/Baskin-Lazpberry/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("baskin")
	fmt.Println(account.Balance(), account.Owner())

	account.Deposit(10)
	fmt.Println(account.Balance(), account.Owner())

	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())

	fmt.Println(account)

}
