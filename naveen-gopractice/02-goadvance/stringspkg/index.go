package stringspkg

import "fmt"  
import "strings"  

func DispIndex() {  
	fmt.Println("The Identifiy To Index")
   str:= "Hi i am in work"  
   fmt.Println(strings.Index(str,"am"))  
}  