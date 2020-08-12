package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "USA"
	re := detectCapitalUse(s)
	fmt.Println(re)
}

func detectCapitalUse(word string) bool {
	if strings.ToLower(word) == word {
		return true
	}

	distance := 'A' - 'a'
	fmt.Println(word[0])
	fmt.Println(word[0] - uint8(distance))
	//word = string(word[0]-uint8(distance)) + word[1:]
	//if strings.ToLower(word) == word {
	//	return true
	//}
	return false
}
