package main

import (
	"fmt"
	"strings"
)

func main() {
	// 输入: 123598765
	// 输出: [123, 987, 876, 765]
	s := "123598765"
	res := getContinuityNaturalNumber(s)
	fmt.Println(res)
	
}

// 获取连续差值为1的自然数
func getContinuityNaturalNumber(s string) []string {
	res := []string{} // [000, 000, 000]
	// byte to slice
	arr := []byte(s)
	for i := 0; i < len(s)-2; i++ {
		// 差值为1的2种可能性
		if arr[i+1]-arr[i] == 1 && arr[i+2]-arr[i+1] == 1 {
			tmp := []string{string(arr[i]), string(arr[i+1]), string(arr[i+2])}
			res = append(res, strings.Join(tmp, ""))
		} else if arr[i]-arr[i+1] == 1 && arr[i+1]-arr[i+2] == 1 {
			tmp := []string{string(arr[i]), string(arr[i+1]), string(arr[i+2])}
			res = append(res, strings.Join(tmp, ""))
		}
	}
	return res
}
