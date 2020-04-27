package functions

import "fmt"

func addret(x int, y int) int {
	fmt.Println("Funtion With Return Value")
	total := 0
	total = x + y
	return total
}