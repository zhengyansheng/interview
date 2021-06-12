package main

import "fmt"

func main() {
	// [0,1,0,0，0，3,12,0]
	arr := []int{0, 1, 0, 3, 0, 0, 12}
	moveZeroesV2(arr)
	fmt.Println(arr)
}

// moveZeroes 原地交换
func moveZeroes(arr []int) {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[j] = arr[i]
			arr[i] = 0
			j++
		}
	}
}

// 双指针
func moveZeroesV2(arr []int) {
	left, right, n := 0, 0, len(arr)
	for right < n {
		if arr[right] != 0 {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
		right++
	}
}
