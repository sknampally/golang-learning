package cases
 
import (
	"fmt"
)
 
func CondCase() {
	fmt.Println("Calling Cond Switch Inti Program")
	var num int =12
	switch {
	case num > 0 && num<10:
		fmt.Println("Single Digit Number")
	case num >10 && num<100:
		fmt.Println("Two Digit Number")
	case num > 99 && num<1000:
		fmt.Println("Three Digit Number")
	case num >999  && num<10000:
		fmt.Println("Four Digit number")
	default:
		fmt.Println("Invalid Data")
	}
	
}