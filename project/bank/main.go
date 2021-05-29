package main

import (
	"fmt"
	"learnGo/project/bank/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
}
