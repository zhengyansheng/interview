package main

import (
	"fmt"
)

/*
	字母排序，给定一个全是小写字母（a,b,c..z）组成的字符串，字母可能重复出现，将字符串按字母升序排列
	例如：输入server，返回eerrsv

https://www.nowcoder.com/discuss/241470
*/

func bubbleSort(arr []rune) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var s = "server"
	s1 := []rune(s)
	fmt.Println(s1)
	bubbleSort(s1)
	fmt.Println(string(s1))
}
