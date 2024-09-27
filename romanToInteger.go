package main

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0

	for i := 0; i < len(s)-1; i++ {
		if romanMap[string(s[i])] >= romanMap[string(s[i+1])] {
			sum += romanMap[string(s[i])]
		} else {
			sum -= romanMap[string(s[i])]
		}
	}

	sum += romanMap[string(s[len(s)-1])]

	return sum
}