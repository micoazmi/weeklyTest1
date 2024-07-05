package main

import "fmt"

func main() {

	fmt.Println(plusOne([]int{1, 2, 3, 4}))
	fmt.Println(plusOne([]int{1, 4, 8, 9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))

}

func plusOne(nums []int) []int {
	result := []int{}
	converts := 0
	for _, num := range nums {
		converts *= 10
		converts += num
	}
	converts += 1

	for converts > 0 {
		i := converts % 10
		result = append([]int{i}, result...)
		converts /= 10
	}
	return result
}
