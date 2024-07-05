package main

import "fmt"

func main() {

	fmt.Println(addDigits([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(addDigits([]int{9, 2, 7}, []int{1, 3, 5}))

}

func addDigits(nums1 []int, nums2 []int) []int {
	convert1 := 0
	convert2 := 0
	for _, num1 := range nums1 {
		convert1 *= 10
		convert1 += num1
	}
	for _, num2 := range nums2 {
		convert2 *= 10
		convert2 += num2
	}
	sum := convert1 + convert2

	result := []int{}
	for sum > 0 {
		digit := sum % 10
		result = append([]int{digit}, result...)
		sum /= 10
	}
	return result

}
