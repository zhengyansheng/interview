# defer

特性
- defer: 执行顺序，栈，先进后出
- return and defer: 先return 在defer
- defer and panic: panic 前的 defer 先执行，然后抛异常
- defer and panic and recovery: panic 前的 defer 出栈先执行，然后抛异常，捕获异常，panic 后的 defer 永远不会执行


## 请说出它对打印结果
```go

package main

import "fmt"

func main() {
	a := 0
	b := 1
	defer cal(a, cal(a, b))
	a = 2
	defer cal(a, cal(a, b))
}

func cal(x, y int) int  {
	fmt.Println(x, y, x+y)
	return x+y
}
```