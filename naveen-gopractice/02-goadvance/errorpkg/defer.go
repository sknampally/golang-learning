package errorpkg
import (  
   "fmt"  
)  
func CallDefer() {  
	fmt.Println("Welcome To Defer Program");
   defer print1("Hello World")  
   print2("Welcome To India")  
}  
func  print1(s string)  {  
   fmt.Println(s)  
}  
func print2(s string)  {  
   fmt.Println(s)  
}  