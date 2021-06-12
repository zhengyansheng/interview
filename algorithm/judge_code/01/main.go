package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	/*
		s.Param = make(map[string]interface{})
		s.Param["name"] = "zhengyansheng"
	*/
	s.Param["name"] = "zhengyansheng"
	fmt.Printf("%#v", s)
}
