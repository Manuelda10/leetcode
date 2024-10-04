package main

import (
	"fmt"
)

func main (){
	strs := "([)]"
	result := isValid(strs)

	fmt.Printf("Es válido: %t", result)
}