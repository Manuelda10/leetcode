package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	revertedNumber := 0
	tmpNum := x

	for tmpNum > 0 {
		revertedNumber = revertedNumber*10 + (tmpNum % 10)
		tmpNum = tmpNum / 10
	}

	return revertedNumber == x
}