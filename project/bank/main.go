package main

import (
	"fmt"
	"learnGo/project/bank/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
