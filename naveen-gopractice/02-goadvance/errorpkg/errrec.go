package errorpkg

import (  
   "fmt"  
)  
func CallRecover() {  
   fmt.Println("Error Recover Program")
   fmt.Println(SaveDivide(10, 0))  
   fmt.Println(SaveDivide(10, 10))  
}  
func SaveDivide(num1, num2 int) int {  
   defer func() {  
      fmt.Println(recover())  
   }()  
   quotient := num1 / num2  
   return quotient  
}  