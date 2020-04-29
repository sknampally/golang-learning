package stringspkg

import "fmt"  
import "strings"  

func DispUpperLower() {  
	fmt.Println("Converting a string to Upper case");
   str := "welcome india"  
   fmt.Println(strings.ToUpper(str))  
   fmt.Println("Converting a string to Lower case");
   str1 := "INDIA HI WORLD"  
   fmt.Println(strings.ToLower(str1))  
}  