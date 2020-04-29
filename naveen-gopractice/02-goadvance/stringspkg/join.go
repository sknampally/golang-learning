package stringspkg

import "fmt"  
import "strings"  

func DispJoin() {  
	fmt.Println("Joining The Strings")
   var arr = []string{"a","b","c","d"}  
   fmt.Println(strings.Join(arr,"*"))  

}  