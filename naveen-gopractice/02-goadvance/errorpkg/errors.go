package errorpkg

import (
	"fmt"
	"errors"
)

func DispErr()	{
	//create errors	
	myErr :=errors.New("Something unexpected happen")
	
	//print error type
	fmt.Println(myErr)
}


	