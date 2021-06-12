package main

import "fmt"

/*
	冒泡排序
*/

func bubbleSort(arr []int) {
	// 函数传参，数组属于 引用传递
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 通过索引原地交换值
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{4, 2, 5, 8, 6, 3, 7, 9, 1}
	fmt.Printf("sort before: %v\n", arr)
	bubbleSort(arr)
	fmt.Printf("sort after: %v\n", arr)

}
