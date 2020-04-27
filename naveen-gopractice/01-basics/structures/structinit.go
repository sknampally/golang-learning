package structures

import "fmt"

type rectangle1 struct {
	length  int
	breadth int
	color   string
}

func DispValues() {
	fmt.Println("Welcome To Struct Intiliaze values")
	var rect1 = rectangle1{11, 12, "Green"}
	fmt.Println(rect1)

	rect2 := new(rectangle1) 
	rect2.length = 21
	rect2.breadth = 22
	rect2.color = "Red"
	fmt.Println(rect2)

	rect3 := rectangle1{length: 31, breadth: 32, color: "Yellow"}
	fmt.Println(rect3)

	
}
