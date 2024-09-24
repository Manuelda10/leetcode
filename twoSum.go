package main

import (
	"fmt"
)

func main (){
	fmt.Printf("Hello, world.\n")
	nums := []int{3,3}
	target := 6


	result := twoSum(nums, target)

	fmt.Printf("La posición de los números es: %v", result)
}

func twoSum(nums []int, target int) []int{

	for i:=0; i<len(nums)-1; i++{
		for j:=i+1; j<len(nums); j++{
			fmt.Printf("Sumó %d y %d \n", nums[i], nums[j])
			tmpSum := nums[i] + nums[j]
			if tmpSum == target {
				return []int{i, j}
			} 
		}
	}

	return nil
}