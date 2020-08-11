package main

import "fmt"

func main() {
	r := findRepeatNumber([]int{1, 2, 3, 3})
	fmt.Println(r)
}

func findRepeatNumber(nums []int) int {
	tmps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := tmps[nums[i]]; ok {
			return nums[i]
		}
		tmps[nums[i]] = 1
	}
	return 0
}
