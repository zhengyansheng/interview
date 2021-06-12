package main

import (
	"fmt"
	"strings"
)

/*

给一个英文句子，词之间有1个或者若干个空格，句子以英文标点“.“结尾，要求颠倒句子中的词序顺序，并且每个词之间只能有一个空格，结尾依旧是”.“，
并且与前一个之间没有空格。

例如输入I     love you  . 返回为you love I.
*/

func reverse(arr []string) ([]string, bool) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr, true
}

func main() {
	var s = "I     love you  ii"
	arr := strings.Fields(s)
	ss, _ := reverse(arr[:len(arr)-1])
	ss = append(ss, ".")
	endS := strings.Join(ss, " ")
	fmt.Println(endS)

}
