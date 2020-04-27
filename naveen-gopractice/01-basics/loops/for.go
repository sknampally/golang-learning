package loops
 
import "fmt"
 
func NumDisp() {
	fmt.Println("Welcome To Loops Program")

	fmt.Println("The Numbers Are");
 	k := 1
	for ; k <= 10; k++ {
		fmt.Print(k,"	")
	}
	println()
	fmt.Println("The Even Numbers Are ")
 	k = 1
	for k <= 10 {
		if	k%2==0 {
				fmt.Print(k,"	")
		}
		k++
	}
	fmt.Println();
	fmt.Println("Prime Numbers ")
	var n int=13
	var c int=0
	for k := 1; ; k++ {
		if n%k==0 {
			c++
		}
		if k == 20 {
			break
		}
	}
	if c==2 {
		fmt.Println("It is prime number")
	} else {
		fmt.Println("Not a prime number")
	}
}