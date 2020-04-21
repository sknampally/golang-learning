package main

import (
	"fmt"
	"strings"
)

var variable string
var strarray []string

func main() {

	
	variable = "Hai Every One Welcomt To GoLang Workspace"

	fmt.Println(variable)

	strarray = strings.Split(variable, " ")

	for i := 0; i < len(strarray); i++ {
		fmt.Println(strarray[i])
	}
}
