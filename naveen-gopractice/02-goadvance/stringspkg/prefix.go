package stringspkg 
 
import "fmt"  
import "strings"  

func DispFix() {  
	fmt.Println("Displaying Prefix & Suffix ")
  
  s := "INDIA"  
   fmt.Println("Prefix :",strings.HasPrefix(s,"IN"))  
   
   s1 := "INDIA"  
   fmt.Println("Suffix :",strings.HasSuffix(s1,"IS"))
}  