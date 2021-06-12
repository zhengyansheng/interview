package main

import (
	"fmt"
)

/*
	我有 1~N 一共 N 个连续的数，每个数出现一次，比如 1，2，3，4，5，6，7，8，9.....N，
	现在我把它打乱，打乱的过程中丢失了一个数，请用时间/空间复杂度尽可能小的方式找出这个数
*/

func main() {
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr := []int{9, 8, 6, 5, 4, 3, 2, 1}
	v, ok := findNum(arr)
	fmt.Println(v, ok)
}

func findNum(arr []int) (int, bool) {
	// []int{1,2,3,4,5,6,7,8,9}
	// []int{9,8,7,6,5,4,3,2,1}
	for i := 0; i < len(arr)-1; i++ {
		// 升序
		if arr[i] < arr[i+1] && arr[i+1]-arr[i] != 1 {
			return arr[i] + 1, true
			// 降序
		} else if arr[i] > arr[i+1] && arr[i]-arr[i+1] != 1 {
			return arr[i] - 1, true
		}
	}
	return 0, false
}
