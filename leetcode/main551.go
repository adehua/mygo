package main

import "fmt"

func main() {
	r := checkRecord("LALL")
	fmt.Println(r)
}

func checkRecord(s string) bool {
	e1, e2 := 0, 0
	exist := true
	for _, v := range s {
		_s := string(v)
		if _s == "A" {
			e1++
			e2 = 0
		} else if _s == "P" {
			e2 = 0
		} else {
			e2++
		}

		if e2 > 2 {
			exist = false
		}
	}

	if exist && e1 <= 1 {
		return true
	}
	return false
}
