package cond
 
import (
	"fmt"
)
 
func DispIf() {
	fmt.Println("Welcome To IF Condition ")
	var s = "Hello World"
	x := true
	if x {
		fmt.Println(s)
	}
	/* we can also use the below syntax for if cond
	if x := 10; x == 10 {
		fmt.Println(s)
	} */
}