package main

import "fmt"

func main() {
	arr := []int{3, 2, 4}
	fmt.Println(twoSum(arr, 6))
	fmt.Println(twoSumV2(arr, 6))
}

func twoSum(numbers []int, target int) []int {
	// write code here
	var res = make([]int, 0)
	if len(numbers) <= 1 {
		return res
	}

	hashMap := make(map[int]int, 0)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, hashMap)
		tValue := target - numbers[i]
		_, ok := hashMap[tValue]
		if ok {
			res = append(res, hashMap[tValue]+1, i+1)
			return res
		} else {
			hashMap[numbers[i]] = i
			// 3 0
			// 2 1
			// 4 2
		}
	}
	return res
}

// 暴力求解
// 时间复杂度 O(n^2)
func twoSumV2(arr []int, target int) []int {
	res := []int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 9 {
				res = append(res, i, j)
				break
			}
		}

	}
	return res
}
