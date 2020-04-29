package stringspkg

import (  
   "fmt"  
   "strings"  
)  

func DispCompare() {  
	fmt.Println("Comparing Two Strings ")
   fmt.Println(strings.Compare("apple", "bat"))  
   fmt.Println(strings.Compare("Hello", "Hello"))  
   fmt.Println(strings.Compare("Dog", "Cat"))  
}  