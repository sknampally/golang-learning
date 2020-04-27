package cases
 
import (
	"fmt"
)
 
func Days() {
	fmt.Println("Calling Switch Program")
	today := 3
 
	switch today {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")	
	default:
		fmt.Println("Sunday")
	}
}
