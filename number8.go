package main

import (
	"fmt"
)

func main() {
	fmt.Println(sliceOperation([]int{1, 2, 3, 4, 5}, []int{2, 4, 6, 7}, 1))
	fmt.Println(sliceOperation([]int{1, 2, 3, 4, 5}, []int{2, 4, 6, 7}, 2))
	fmt.Println(sliceOperation([]int{1, 2, 3, 4, 5}, []int{2, 4, 6, 7}, 3))
	fmt.Println(sliceOperation([]int{1, 2, 3, 4, 5}, []int{2, 4, 6, 7}, 4))

}

func sliceOperation(nums1 []int, nums2 []int, typeOperation int) []int {
	results := []int{}

	switch typeOperation {
	case 1:
		for _, num1 := range nums1 {
			for _, num2 := range nums2 {
				if num1 != num2 {
					if !found(results, num1) {
						results = append(results, num1)

					}
				}
			}
		}
	case 2:
		for _, num2 := range nums2 {
			for _, num1 := range nums1 {
				if num2 != num1 {
					if !found(results, num2) {
						results = append(results, num2)
					}
				}
			}
		}

	case 3:
		for _, num1 := range nums1 {
			results = append(results, num1)
		}

		for _, num2 := range nums2 {
			if !found(results, num2) {
				results = append(results, num2)

			}
		}
	case 4:
		for _, num1 := range nums1 {
			for _, num2 := range nums2 {
				if num1 == num2 {
					if !found(results, num2) {
						results = append(results, num2)
					}
				}
			}
		}

	}

	return results

}

func found(res []int, num1 int) bool {
	for _, v := range res {
		if num1 == v {
			return true
		}
	}
	return false
}
