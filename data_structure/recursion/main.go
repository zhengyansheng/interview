package main

import (
	"strings"
)

func main()  {
	
}

// StrToUpper 字符串全部转大写
func StrToUpper(s string) []string {
	rs := []string{}

	var fn func(int, []byte)
	fn = func(i int, bytes []byte) {
		// 0 7
		if i == len(bytes)-1 {
			return
		}
		rs = append(rs, strings.ToUpper(string(bytes[i])))
		i++
		fn(i, bytes)
	}
	fn(0, []byte(s))
	return rs
}
