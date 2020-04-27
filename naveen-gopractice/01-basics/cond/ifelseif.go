package cond
 
 import (
	"fmt"
)
 
func Bigger() {
	fmt.Println("Welcome To If Else If Program ")
	x := 100
	y:= 50
	z:= 80
 
	if x > y && x > z {
		fmt.Println("X is Bigger")
	} else if y > x && y > z {
		fmt.Println("Y is Bigger")
	} else {
		fmt.Println("Z is Bigger")
	}
}