package stringspkg

import "fmt"  
import "strings"  

func DispCount() {  
	fmt.Println("Displaying the Count ")
   str:= "I am an indian welcome to hyderabad"  
   fmt.Println(strings.Count(str,"e"))  
}  