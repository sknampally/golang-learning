package main

import (
	"fmt"
)
func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}


func main() {
	fmt.Println(plus(3, minus(10, 5)))
}
