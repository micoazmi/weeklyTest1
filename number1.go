package main

import "fmt"

func main() {
	fmt.Println(removeDuplicateLetter("bananas"))
	fmt.Println(removeDuplicateLetter("lalalamama"))
	fmt.Println(removeDuplicateLetter("abcbcade"))
}

func removeDuplicateLetter(words string) string {
	duplicate := ""
	for i := 0; i < len(words); i++ {

		if !found(words[i], duplicate) {
			duplicate += string(words[i])
		}

	}

	return duplicate
}

func found(word byte, duplicates string) bool {
	for i := 0; i < len(duplicates); i++ {
		if duplicates[i] == word {
			return true
		}
	}
	return false

}
