package pointers

import (  
    "fmt"
)

func hello() *int {  
    i := 5
    return &i
}
func retval() {  
     fmt.Println("Returing a  pointer from function")
    d := hello()
    fmt.Println("Value of d", *d)
}