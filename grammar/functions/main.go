package main

import (
	"fmt"
	"strings"
)

// naked return

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLenght, _ := lenAndUpper("nico")
	fmt.Println(totalLenght)

}
