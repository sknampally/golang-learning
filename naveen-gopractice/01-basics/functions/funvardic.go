package functions
 
import "fmt"
 
func variadicExample(s ...string) {
	fmt.Println("Varadic Funtions Values Are")
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
    fmt.Println(s[3])
	fmt.Println(s[4])
}