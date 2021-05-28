package main

import "fmt"

// map[key]value{key: value}

func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}
}
