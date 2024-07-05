package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(OddBeforeEven(number))
}

func OddBeforeEven(nums []int) []int {
	results := []int{}

	for _, num := range nums {
		if num%2 != 0 {
			results = append(results, num)
		}
	}

	for _, num := range nums {
		if num%2 == 0 {
			results = append(results, num)
		}
	}
	return results

}
