package main

import "fmt"

type student struct {
	Name string
}

func isOk(v interface{}) {
	switch msg := v.(type) {
	//case *student, student: 错误的写法
	case *student:
		fmt.Println(msg.Name)
	case student:
		fmt.Println(msg.Name)
	}
}

func main() {
	var s student
	s.Name = "zhengyansheng"
	isOk(s)
}
