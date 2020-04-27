package functions

import "fmt"

func add(x int, y int) {
	fmt.Println("Simple Addition Function");
	total := 0
	total = x + y
	fmt.Println("The Sum Is :",total)
}
