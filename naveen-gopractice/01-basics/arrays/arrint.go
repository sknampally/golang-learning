package arrays

import "fmt"

func DispArrInit() {
	x := [5]int{10, 20, 30, 40, 50}   
	var y [5]int = [5]int{10, 20, 30} 
	fmt.Println("The Values of Array X");
	fmt.Println(x)
	fmt.Println("The Values of Array Y");
	fmt.Println(y)
}