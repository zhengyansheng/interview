package main

import "fmt"

/*
	请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
	给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

*/

func reverseString(s string) (string, bool) {
	length := len([]byte(s))
	if length > 5000 {
		return fmt.Sprintf("%s length more than 5000.", s), false
	}
	s2 := []byte(s)
	for i := 0; i < length/2; i++ {
		s2[i], s2[length-i-1] = s2[length-i-1], s2[i]
	}
	return string(s2), true
}

func main() {
	//s := "abcde"
	//s := "abcd"
	s := "a"
	fmt.Println(reverseString(s))
}
