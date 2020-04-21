package main

import (
	"fmt"
)

var strarray = []string{"lorem", "ipsum", "dolor", "sit", "amet"}
var intarray = []int{1, 2, 4, 8, 16}
var mapone = map[int]string{}
var maptwo = map[string]interface{}{}

func main() {

	for i := 0; i != 5; i++ {

		fmt.Println(intarray[i], "\t", strarray[i])

		mapone[intarray[i]] = strarray[i]
		maptwo[strarray[i]] = mapone
	}
	fmt.Println(mapone)
	fmt.Println(maptwo)
}
