package arrays

import "fmt"

func DispArrValues() {
	fmt.Println("The Arrays Values Are");
	var theArray [3]string
	theArray[0] = "India"  
	theArray[1] = "Canada" 
	theArray[2] = "Japan"  

	fmt.Println(theArray[0]) 
	fmt.Println(theArray[1]) 
	fmt.Println(theArray[2]) 
}