package errorpkg

import (
	"fmt"
)

func CallErrors()	{
	fmt.Println("Calling Error Handling Programs ")
	 //Calling Error throwing program
	DispErr()
	//Calling Divide by zero program
	CallDiv()
	//Calling Invalid Number program
	DispInvalidNum()
	//Calling Error File program
	CallFile()
	//Calling Error recover program
	CallRecover()
	//Calling Error Defer program
	CallDefer()
	//Calling Error Panic program
	CallPanic()
}


	