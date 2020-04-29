package main

import (
	"fmt"
	e "errorpkg"
	st "stringspkg"
)
func main() {
	fmt.Println("Welcome To Advance Programing")
	 //Calling Error Handilng
	e.CallErrors()
	//Calling String Handling
	st.CallFun()

}

/*
import (
	"fmt"
	d "datatypes"
	c "cond"
	s "cases"
	l "loops"
	f "functions"
	st "structures"
	ar "arrays"
    p "pointers"
)
func main() {

	fmt.Println("Welcome To GoLang Basics Programs")
	var input int
	for
	{
		fmt.Println("/-------------------------------------/");
		fmt.Println("	1.DataTypes	");
		fmt.Println("	2.Conditions(IF/IF--ELSE)	");
		fmt.Println("	3.Switch Cae	");
		fmt.Println("	4.Loops	");
		fmt.Println("	5.Arrays	");
		fmt.Println("	6.Functions	");
		fmt.Println("	7.Pointres	");
		fmt.Println("	8.Structures	");
		fmt.Println("	9.Exit	");
		fmt.Println("/------------------------------------/");
		fmt.Print("Enter the topic : ")
		fmt.Scanf("%d", &input)
		if input==1 {
			d.CallDatatypes()
		} else if input==2 {
			c.CallConditions()
		} else if input==3 {
			s.CallCase()
		} else if input==4 {
			l.CallLoops()
		} else if input==5 {
			ar.CallArrays()
		} else if input==6 {
			f.CallFunc()
		} else if input==7 {
			p.CallPointers() 
		} else if input==8 {
			st.CallStruct()
		} else if input==9 {
			break
		} else {
			fmt.Println("Invalid Option ")
		}
	}
}*/