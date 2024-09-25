package main

import (
	"fmt"
)

func main (){
	number := 1991
	result := isPalindrome(number)

	fmt.Printf("El número %d es palíndromo: %t", number, result)
}