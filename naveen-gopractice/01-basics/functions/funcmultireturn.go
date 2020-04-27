package functions

import "fmt"

func rectanglemulti(l int, b int) (area int, parameter int) {
	fmt.Println("Funtion With Named Return Value")
	parameter = 2 * (l + b)
	area = l * b
	return
}
