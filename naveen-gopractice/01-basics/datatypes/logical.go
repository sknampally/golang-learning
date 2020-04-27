package datatypes

import "fmt"

func LogicalDisp() {
	fmt.Println("The Logical Operators ")
	var x, y, z = 10, 20, 30

	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x == y && x > z))
}