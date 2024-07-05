package main

import "fmt"

func main() {
	fmt.Println(sameElements([]int{1, 2, 4, 7, 8}, []int{2, 3, 7, 9}))
	fmt.Println(sameElements([]int{1, 2, 7, 4, 7, 8}, []int{7, 7, 3, 2, 9}))
	fmt.Println(sameElements([]int{2, 4, 6, 8}, []int{1, 3, 5, 7, 9}))

}

func sameElements(nums1 []int, nums2 []int) []int {
	result := []int{}

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if num1 == num2 {
				if !found(result, num2) {
					result = append(result, num2)

				}
			}
		}
	}

	return result
}

func found(res []int, num2 int) bool {
	for _, v := range res {
		if num2 == v {
			return true
		}
	}
	return false
}
