package arrays

import (
	"fmt"
	"reflect"
)

func DispArrEllipse() {

	x := [...]int{10, 20, 30,40}


	fmt.Println("The Type Of Array is :",reflect.ValueOf(x).Kind())
	fmt.Println("The Lenght is :",len(x))

}