package main

import "fmt"

type People struct {
	Name string
}

// 产生了循环引用
// fmt.Sprintf("print: %v", p) 打印p时 就是调用了 p的String方法
func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	fmt.Println(p)
}
