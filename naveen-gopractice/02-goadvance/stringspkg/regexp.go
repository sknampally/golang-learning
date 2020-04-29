package stringspkg	
  
import (  
    "fmt"  
    "regexp"  
)  
  
func DispRegExp() {  
	fmt.Println("The FindString Returns Left value")
    re := regexp.MustCompile(".com")  
    fmt.Println(re.FindString("google.com"))  
    fmt.Println(re.FindString("abc.org"))  
    fmt.Println(re.FindString("fb.com"))  
	fmt.Println("The FindStringIndex Returns Index Value")
    re1 := regexp.MustCompile(".com")  
    fmt.Println(re1.FindStringIndex("google.com"))  
    fmt.Println(re1.FindStringIndex("abc.org"))  
    fmt.Println(re1.FindStringIndex("fb.com"))  
}  