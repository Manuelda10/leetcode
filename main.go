package main

import (
	"fmt"
)

func main (){
	strs := []string{"ca"}
	result := longestCommonPrefix(strs)

	fmt.Printf("El prefijo común es: %s", result)
}