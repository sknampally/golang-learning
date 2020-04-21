package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	mySlice = append(mySlice[:4], mySlice[5:]...)
	fmt.Println(mySlice)
}
