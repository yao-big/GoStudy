package main

import "fmt"

/*
1. 数组
	使用一种数据类型的固定长度的序列
	数组定义: var a [len]int 比如: var a[5]int 代表长度为5的数组 数组长度必须为常量, 且是类型的组成部分, 一旦定义, 长度不能改变
	长度是数组类型的一部分, 因此 var a[5]int 和 var a[10]int 是不同的类型
	数组的长度必须是常量, 表达式是不允许的, 比如 array := [3.5]int{} 是不允许的
	数组可以通过下标进行访问, 下标是从 0 开始, 最后一个元素的下标是 len(a)-1

    for i:=0;i<len(a); i++ {
		fmt.Println(a[i])
	}
	// 也可以使用 range 迭代数组元
	for index,_ := range a {
		fmt.Println(a[index])
	}

    访问越界 , 如果下标现在数组范围之外, 则触发访问越界 会 panic
	数组是值类型, 赋值和传值会复制整个数组, 而不是指针, 因此会改变副本的值, 不会改变本身的值
	支持 == != 操作符, 因为内存总是被初始化的,
	指针数组 [n]*T 数组指针 *[n]T


内置函数 len 和 cap 可以赶回数组的长度和容量

*/

func main() {
	// 数组初始化:
	var arr0 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr0)
	// 打印类型
	fmt.Printf("%T\n", arr0)

	b := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(len(b))

	d := [...]struct {
		name string
		age  int
	}{
		{"name", 1},
		{"name", 2},
	}
	fmt.Println(d)

	var str = [5]string{3: "hello world", 4: "tom"}
	fmt.Println(len(str))

	//遍历多维数组
	var f [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}

	// k1 是索引, v1 是值
	for k1, v1 := range f {
		fmt.Println(k1, v1)
		//for k2, v2 := range v1 {
		//	fmt.Printf("f[%d][%d]=%d\n", k1, k2, v2)
		//}
	}

	alist := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, i := range alist {
		sum += i
	}

	fmt.Println(sum)

	fmt.Println("=====================================")
	// 找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
	// 找出两个元素之和等于8的下标分别是（0，4）和（1，2）

	alist2 := [5]int{1, 3, 5, 8, 7}
	myTest(alist2, 8)

}

func myTest(alist [5]int, target int) {
	for i, _ := range alist {
		other := target - alist[i]
		fmt.Println(other)
		for j := i + 1; j < len(alist); j++ {
			if alist[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
