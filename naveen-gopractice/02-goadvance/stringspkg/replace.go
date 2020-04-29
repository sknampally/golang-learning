package stringspkg

import "fmt"  
import "strings"  

func DispReplace() {  
	fmt.Println("Replacing a letters")
   str:= "Welcome to world what is going on"  
   fmt.Println(strings.Replace(str,"w","Z",2))  
}  