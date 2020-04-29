package stringspkg

import "fmt"  
import "strings"  

func DispSplit() {  
	fmt.Println("Displaying the values using Split Function")
   str := "I,love,my,country"  
   var arr []string = strings.Split(str, ",")  
   fmt.Println(len(arr))  
   for i, v := range arr {  
      fmt.Println("Index : ", i, "value : ", v)  
   }  
}  