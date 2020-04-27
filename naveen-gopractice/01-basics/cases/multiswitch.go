package cases
 
import (
	"fmt"
)
 
func MultiCase() {
	fmt.Println("Calling Multi Switch Program")
	var t int = 6
 
	switch t {
	case 2, 4, 6, 8:
		fmt.Println("Even Number")
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd Number.")
	default:
		fmt.Println("Not a Valid Number")
	}
}