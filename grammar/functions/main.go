package main

import (
	"fmt"
	"strings"
)

// naked return

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done") // defer : after func run
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLenght, up := lenAndUpper("nico")
	fmt.Println(totalLenght, up)

}
