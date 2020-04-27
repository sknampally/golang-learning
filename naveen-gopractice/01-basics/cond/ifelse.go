package cond
 
import (
	"fmt"
)
 
func Even() {
	fmt.Println("Welcome To If Else Program ")
	x := 100
 
	if x%2 == 0 {
		fmt.Println("It is Even Number")
	} else {
		fmt.Println("It is Odd Number")
	}
}