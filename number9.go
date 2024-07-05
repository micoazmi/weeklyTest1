package main

import "fmt"

func main() {
	number1 := []int{1, 2, 3, 4, 4, 4, 3, 3, 2, 4}
	number2 := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	fmt.Println(countFrequentElement(number1))
	fmt.Println(countFrequentElement(number2))
}

func countFrequentElement(nums []int) map[int]int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num]++
	}
	return result
}
