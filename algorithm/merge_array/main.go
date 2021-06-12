package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8, 9, 10}
	arr := mergeArray2(arr1, arr2)
	fmt.Println(arr)
}

func mergeArray(arr1, arr2 []int) []int {
	// 1. 合并
	for _, v := range arr1 {
		arr2 = append(arr2, v)
	}

	// 2. 升序
	sort.Ints(arr2)

	// 3. return
	return arr2
}

func mergeArray2(arr1, arr2 []int) []int {
	arr := []int{}

	isContinue := len(arr1) > len(arr2)
	for i := 0; i < len(arr1); i++ {
		if isContinue && i == len(arr2) {
			continue
		}
		if arr1[i] < arr2[i] {
			arr = append(arr, []int{arr1[i], arr2[i]}...)
		} else if arr1[i] > arr2[i] {
			arr = append(arr, []int{arr2[i], arr1[i]}...)
		} else {
			arr = append(arr, []int{arr1[i], arr2[i]}...)
		}
	}
	if isContinue {
		arr = append(arr, arr1[len(arr2):]...)
	} else {
		arr = append(arr, arr2[len(arr1):]...)
	}

	return arr
}

