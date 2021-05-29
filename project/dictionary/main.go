package main

import (
	"fmt"
	"learnGo/project/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
}
