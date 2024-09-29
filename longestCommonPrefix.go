package main

func longestCommonPrefix(strs []string) string {
	// Sacar todas las posibles combinaciones de la primera palabra. Empezar comparando la más larga hacia la más corta en cada palabra.
	// Tomar como referencia la palabra más corta
	minIndex := 0
	minWord := strs[0]

	for i := 1; i < len(strs); i++ {
		if len(minWord) > len(strs[i]) {
			minWord = strs[i]
			minIndex = i
		}
	}

	//Intercambiar posiciones del mínimo y el primero para que el minimo esté al inicio.
	tmp := strs[0]
	strs[0] = minWord
	strs[minIndex] = tmp

	longestPrefix := ""

	//Recorrer el array tomando como base la palabra más corta como máximo prefijo, y luego ir reduciendo el tamaño del prefijo hasta 0
	for j := len(strs[0]); j > 0; j-- {
		longestPrefix = strs[0][0:j]

		// Que empiece en 0 y se compare con sí misma
		for i := 1; i < len(strs); i++ {
			if longestPrefix == strs[i][0:j] {
				continue
			}
			longestPrefix = ""
			break
		}

		if longestPrefix != "" {
			return longestPrefix
		}
	}

	return longestPrefix
}