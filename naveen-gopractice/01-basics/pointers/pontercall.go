package pointers

import (  
    "fmt"
)

func change(val *int) {  
    *val = 55
}
func swapdisp() {  
    fmt.Println("The Value Change Function")
    a := 58
    fmt.Println("value of a before function call is",a)
    b := &a
    change(b)
    fmt.Println("value of a after function call is", a)
}