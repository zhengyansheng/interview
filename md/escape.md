# 逃逸分析

## 1. 什么是逃逸

> 编译器逃逸分析后 最终控制。



人很难区分一个变量是要分配到堆上还是分配到栈上，因为这都是Go编译器来控制的。



> 1. 如果函数外部没有引用，则优先放到栈 Stack中；
>
> 2. 如果函数外部存在引用，则必定放到堆 Heap中；



```go
package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(*x)

}

func foo() *int {
	t := 1
  fmt.Println(t) // ...
	return &t
}
```



```bash
✗ go build -gcflags '-l -m' main.go
# command-line-arguments
./main.go:12:2: moved to heap: t
./main.go:13:13: ... argument does not escape
./main.go:13:13: t escapes to heap  // t原本是要分配到栈上到，但是由于函数的返回值有被外部调用，强制被编译器分配到了堆上。
./main.go:7:13: ... argument does not escape
./main.go:7:14: *x escapes to heap
```

