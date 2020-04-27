package loops
 
import "fmt"
 
func RangeDisp() {
	fmt.Println("The For range Display")
	var str="Hello"
	for range str {
		fmt.Println("Hello")
	}
}