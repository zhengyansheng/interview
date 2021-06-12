package main

import "fmt"

func main() {
	// [1, 2, 3, 4, 5, 7, 9] 5
	arr := []int{1, 2, 3, 4, 5, 7, 9}
	target := 5
	idx := binarySearch(arr, target)
	fmt.Println(idx)
}

// 二分查找模版
// 时间复杂度: O(LogN)
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
