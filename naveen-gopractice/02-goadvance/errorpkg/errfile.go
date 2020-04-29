package errorpkg

import (  
    "fmt"
    "os"
)

func CallFile() {  
	fmt.Println("Welcome To Open File Program")
    f, err := os.Open("/test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}