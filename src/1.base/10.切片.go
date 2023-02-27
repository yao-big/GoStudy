package main

import "fmt"

/*

1. 切片
	slice 不是数组或者是数组指针, 他通过内部指针和相关属性引用数组片段, 以实现变长方案

*/
func main() {
	// hello.Print()
	//编译报错：./main.go:6:5: undefined: hell
	fmt.Println("hello world")
}
