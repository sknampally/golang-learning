package pointers

import (  
    "fmt"
)

func DispVal() {  
    b := 255
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
    fmt.Println("address of b is", a)
	fmt.Println("Values of b is:",b)
}