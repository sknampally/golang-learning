package main

import (
	"fmt"
)

func named() (str string) {
	str = "hello"
	return
}

func typed() string {
	var str = "world"
	return str
}


func main() {
	fmt.Println(named())
	fmt.Println(typed())
}
