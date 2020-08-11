package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcabc"
	re := repeatedSubstringPattern(s)
	fmt.Println(re)
}

func repeatedSubstringPattern(s string) bool {
	var str1 string = s + s
	var str2 string = str1[1 : len(str1)-1]
	if strings.Contains(str2, s) {
		return true
	} else {
		return false
	}
}
