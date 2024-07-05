package main

import "fmt"

func main() {
	number := []int{1, 2, 2, 3, 3, 4, 4, 4}
	fmt.Println(removeDuplicate(number))

}

func removeDuplicate(nums []int) []int {
	results := []int{}

	for _, num := range nums {
		found := false
		if found {
			results = append(results, num)
		}

	}
	return results
}
