package main

func isValid(s string) bool {
	// Diccionario que mapea los paréntesis de cierre con sus correspondientes de apertura
	matchingBrackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Pila para almacenar los paréntesis de apertura
	stack := []rune{}

	// Recorremos cada carácter de la cadena
	for _, char := range s {
		// Si el carácter es un paréntesis de cierre
		if open, found := matchingBrackets[char]; found {
			// Verificamos si la pila está vacía o si el último elemento no coincide
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// Quitamos el último elemento de la pila
			stack = stack[:len(stack)-1]
		} else {
			// Si es un paréntesis de apertura, lo añadimos a la pila
			stack = append(stack, char)
		}
	}

	// La cadena es válida si la pila está vacía al final
	return len(stack) == 0
}