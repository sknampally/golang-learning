package functions

import "fmt"

func rectangle(l int, b int) (area int) {
	fmt.Println("Funtion With Named Return Value")
	
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)
	area = l * b
	return 
}

