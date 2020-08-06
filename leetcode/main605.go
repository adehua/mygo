package main

import "fmt"

func main() {
	flowerbed := []int{1,0,0,0,0,0,1}
	n := 2
	re := canPlaceFlowers(flowerbed, n)
	fmt.Println(re)

}

func canPlaceFlowers(flowerbed []int, n int) bool {
	len := len(flowerbed)
	for k, v := range flowerbed {
		if v == 0 && (k == 0 || (k > 0 && flowerbed[k-1] == 0)) && (k == len-1 || (k != len-1 && flowerbed[k+1] == 0)) {
			flowerbed[k] = 1
			n--
		}
	}

	if n <= 0 {
		return true
	} else {
		return false
	}
}
