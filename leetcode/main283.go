package main

import "fmt"

func main() {
	r := moveZeroes([]int{0, 1, 0, 3, 12})
	fmt.Println(r)
}
func moveZeroes(nums []int) int {
	len := len(nums)
	l := 0
	z := 0
	for i := 0; i < len; i++ {
		num := nums[i]
		if num == 0 {
			z++
		} else {
			nums[l] = nums[i]
			l++
		}
	}
	for i := len - 1; i >= len-z; i-- {
		nums[i] = 0
	}
	fmt.Println(nums)
	return 0
}
