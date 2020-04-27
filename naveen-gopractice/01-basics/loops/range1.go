package loops
 
import "fmt"
 
func Range1Disp() {
	
	fmt.Println("The Range Values Are");
	a := []string{"Hello", "World","Welcome","Fine","Bye"}
	for i, s := range a {
		fmt.Println("Index : ",i,"Value : ", s)
	}
}