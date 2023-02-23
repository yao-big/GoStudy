package main

import "fmt"

/*
init 函数用于包的初始化,
特征:
	1. 用于程序执行前做包的初始化函数, 比如初始化包里的变量等,
	2. 每个包可以有多个 init 函数
	3. 包的每个源文件也可以有多个 init 函数,
	4. 同一个包中多个 init 函数的执行顺序在 go 没有明确的定义
	5. 不同包的 init 函数按照包导入的依赖关系决定执行顺序
	6. init 函数不能被其他函数调用, 而在 main 函数之前被调用


相同点:
	两个函数定时不能有任何的参数和返回值, 且 都是 go 自动调用
不同点:
	init 可以在任意包中, 且可以重复定义多个
	main 函数只能用于 main 包中, 且只能定义一个

对同一个文件的 init 调用顺序是从上到下的
对同一个 package 中不同文件 按照文件名 比较 从小到大的顺序调用 init 函数
对不同的 package  如果不互相依赖的话, 按照 main 包中的 先 import 的后调用 其包中的 init() 如果 package 存在依赖, 则先调用最早被依赖的 package 中的 init 函数 最后再用 main 函数

如果 init 中使用了 println 或者 print 不会按照顺序来, 这两个函数只推荐在 测试环境中使用, 对于正式环境不要使用
*/

func init() {
	fmt.Println("首先执行了 init 函数")
}

func main() {
	fmt.Println("main 函数")
}
