package cases
 
import (
	"fmt"
)
 
func FallCase() {

	fmt.Println("Calling Fall Switch Program")
	day := 2
 
	switch day {
	case 1:
		fmt.Println("Monday ")
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
		fallthrough
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("No information available.")
	}
}