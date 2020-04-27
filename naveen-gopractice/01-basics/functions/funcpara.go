package functions
 
import "fmt"
 
func calculation(str string, y ...int) int {
	fmt.Println("Functional Parameter with Varadic Funtions ")
	fmt.Println("The Values are");
 
    area := 1
 
    for _, val := range y {
        if str == "Rectangle" {
            area *= val
        } else if str == "Square" {
            area = val * val
        }
    }
    return area
}