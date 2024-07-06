package main

import (
	"fmt"
)

func main() {
	fruits1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	fruits2 := []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}
	fmt.Println(sliceFruits(fruits1, fruits2, 1))
	fmt.Println(sliceFruits(fruits1, fruits2, 2))
	fmt.Println(sliceFruits(fruits1, fruits2, 3))

}

func sliceFruits(fruits1 []string, fruits2 []string, operationType int) []string {
	defaultAnswer := []string{"please input type operation between 1 or 2"}
	result := []string{}
	switch operationType {
	case 1:
		for _, fruit1 := range fruits1 {
			if found(fruits2, fruit1) {
				if !found(result, fruit1) {
					result = append(result, fruit1)

				}
			}
		}
	case 2:
		for _, fruit1 := range fruits1 {
			if !found(fruits2, fruit1) {
				if !found(result, fruit1) {
					result = append(result, fruit1)

				}
			}
		}
		for _, fruit2 := range fruits2 {
			if !found(fruits1, fruit2) {
				result = append(result, fruit2)
			}
		}
	default:
		return defaultAnswer
	}

	return result

}

func found(res []string, fruit2 string) bool {
	for _, v := range res {
		if fruit2 == v {
			return true
		}
	}
	return false

}
