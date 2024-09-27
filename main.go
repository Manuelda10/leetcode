package main

import (
	"fmt"
)

func main (){
	number := "XIX"
	result := romanToInt(number)

	fmt.Printf("El número %s en entero es: %d", number, result)
}