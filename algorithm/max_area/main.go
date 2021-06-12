package main

import (
	"fmt"
)

func main() {
	// [1,8,6,2,5,4,8,3,7]
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	maxAreaValue := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			tmpMaxArea := (j - i) * min(height[i], height[j])
			maxAreaValue = max(maxAreaValue, tmpMaxArea)
		}
	}
	return maxAreaValue
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
