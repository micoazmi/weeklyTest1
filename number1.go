package main

import "fmt"

func main() {
	fmt.Println(removeDuplicateLetter("bananas"))
	fmt.Println(removeDuplicateLetter("lalalamama"))

}

func removeDuplicateLetter(words string) string {
	duplicate := ""
	for i := 0; i < len(words); i++ {
		found := false

		for j := 0; j < len(duplicate); j++ {
			if duplicate[j] == words[i] {
				found = true
				break
			}
		}
		if !found {
			duplicate += string(words[i])
		}

	}

	return duplicate
}
