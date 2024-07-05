package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isAnagram("Otto", "Toto"))
	fmt.Println(isAnagram("Ayam", "Maya"))
	fmt.Println(isAnagram("Tamat", "Kiamat"))
	fmt.Println(isAnagram("Tamat", "matat"))

}

func isAnagram(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	mapWord1 := make(map[rune]int)
	mapWord2 := make(map[rune]int)

	for _, v := range word1 {
		mapWord1[v]++
	}

	for _, v := range word2 {
		mapWord2[v]++
	}

	for i, v := range mapWord1 {
		if mapWord2[i] != v {
			return false
		}
	}
	return true

}
