package main

import "fmt"

/*
	快速排序
	- 从数组中取一个基准值，默认取索引为0的第一个值，假设值为 mid
	- 用 mid 分别和数组从右往左进行值比较，如果比 mid 值大，就交换位置，
	- 再次用 mid的值 和数组从左往右进行值比较，如果比 mid 值大，就交换位置，
	- 以此类推，最终就得到了结果
*/

func quickSort(arr []int, start, end int) {
	// 从小到大排序
	if start >= end {
		return
	}
	// 4
	mid := arr[start]
	// 0, 8
	i, j := start, end
	for i < j {
		for i < j && mid < arr[j] {
			j--
		}
		if i < j {
			arr[j] = arr[i]
			i++
		}

		for i < j && mid > arr[i] {
			i++
		}
		if i < j {
			arr[i] = arr[j]
			j--
		}
	}
	arr[i] = mid
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

func main() {
	arr := []int{4, 2, 5, 8, 6, 3, 1, 9, 7}
	fmt.Printf("sort before: %v\n", arr)
	// {4, 2, 5, 8, 6, 3, 1, 9, 7}, 0, 9
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("sort after: %v\n", arr)

}
