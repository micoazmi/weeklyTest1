package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := []string{"this", "is", "a", "kanoha"}
	word2 := []string{"is", "a"}

	fmt.Println(capitalize(word1, word2))
}

func capitalize(words []string, except []string) []string {
	result := []string{}

	for _, word := range words {
		if contain(word, except) {
			result = append(result, word)
		} else {
			result = append(result, strings.Title(word))
		}
	}
	return result

}

func contain(word string, exceptWords []string) bool {
	for _, exceptWord := range exceptWords {
		if word == exceptWord {
			return true
		}
	}
	return false

}
