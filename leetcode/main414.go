package main

import (
	"fmt"
	"math"
)

func main() {
	r := thirdMax([]int{3,2,1})
	fmt.Println(r)
}
func thirdMax(nums []int) int {
	var firstMax int
	preMax := math.MaxInt32
	var exist bool
	for i := 0; i < 3; i++ {
		max := math.MinInt32
		exist = false
		for _, n := range nums {
			if n < preMax {
				exist = true
				if n > max {
					max = n
				}
			}
		}
		if i == 0 {
			firstMax = max
		}
		if !exist {
			break
		}
		preMax = max
	}

	if !exist {
		return firstMax
	}
	return preMax
}
