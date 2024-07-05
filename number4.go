package main

import "fmt"

func main() {
	number := []int{1, 2, 2, 3, 3, 4, 4, 4}
	fmt.Println(removeDuplicate(number))

}

func removeDuplicate(nums []int) []int {
	results := []int{}

	for i := 0; i < len(nums); i++ {
		if !found(results, nums[i]) {
			results = append(results, nums[i])
		}
	}
	return results
}

func found(results []int, num int) bool {
	for _, result := range results {
		if result == num {
			return true
		}
	}
	return false
}
