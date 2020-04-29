package errorpkg

import (
	"fmt"
	"errors"
)

func DivideErr(a int,b int)	(int, error) { 
	if b==0 {
		return 0,errors.New("Cannot be divide by Zero")
	} else {
		return (a/b) ,nil
	}
}
func CallDiv() {
	fmt.Println("Welcome To Divide By Zero Program")
	if result, err:=DivideErr(4,2);err!=nil {
		fmt.Println("Error Occured",err)
	} else {
		fmt.Println("4/0 is ",result)
	}
}


	