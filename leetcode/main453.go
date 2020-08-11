package main

import (
	"fmt"
	"sort"
)

func main() {
	r := minMoves([]int{1, 2, 4, 3})
	fmt.Println(r)
}

func minMoves(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}
	sort.Ints(nums)

	count := 0
	for i := 1; i < l; i++ {
		count += nums[i] - nums[0]
	}
	return count
}
