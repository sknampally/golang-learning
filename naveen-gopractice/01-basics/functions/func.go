package functions

import (
	"fmt"
)
func CallFunc() {
	fmt.Println("Welcome To Functions Programs ")
	//Calling a Simple Functioj
	add(20,30)
	//Calling a function with return values
	x:=addret(50,40);
	fmt.Println("The Return Sum Is :",x);
	//Calling a function with named return
	area:=rectangle(20, 30)
	fmt.Println("The Area is :",area);
	//Calling a function with multiple return values
	var a, p int
	a, p = rectanglemulti(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
	//Calling a Variadic function 
	variadicExample("Hello", "World", "Welcome", "Fine","Bye")
	//Calling a Functional Parameetrs with Variadic function 
	fmt.Println(calculation("Rectangle", 20, 30))
    fmt.Println(calculation("Square", 20))
	
}
	
