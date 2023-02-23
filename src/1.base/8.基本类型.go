package main

import (
	"fmt"
	_ "fmt"
	"math"
	"strings"
)

/*
1. 整型:
	1. int8 ,int16,int32 ,int64
	2. uint8 , uint16,uint32,uint64

	uint8 就是 byte 类型,

2. 浮点类型
	float32
	float64

3. 复数:
	实部:real
	虚部:imag

4. 布尔值:
	布尔类型只有 true 和 false
	默认值为 false
	go 语言不允许将整型强制转换为布尔类型
	布尔类型无法参与运算, 也无法和其他类型转换

5. 字符串
	s1 := "hello "

6. 字符串转义符
	常见转义符包括 回车, 换行, 单双引号, 制表符等
	\r 回车符
	\n 换行符
	\t 制表符
	\' 单引号
	\" 双引号
	\  反斜杠

7. 多行字符串
	使用反引号字符,
	反引号字符可以包含任何内容, 包括换行符

8.字符串的常用操作
	len() 获取字符串的长度
	+ 或者 fmt.Sprintf  拼接字符串
	strings.split() 分割
	strings.Contains() 判断是否包含
	strings.HasPrefix() 前缀是否包含
	strings.HasSuffix() 后缀是否包含

9. byte 和 rune 类型
	组成每个字符串的元素叫字符, 可以通过遍历或者单个原货字符串元素来获得字符, 字符用单引号括起来

	unit8 类型, 或者叫 byte 类型 代表了一个 ascii 代码
	rune 类型代表了一个 utf-8 类型

	当需要处理中文和日本 其他复杂类型的字符时 , 就需要用到 rune 类型, rune 类型实际是一个 int 32 类型,

10. 修改字符串
	1. 修改字符串, 需要先将其转换成 []rune 或者 []byte 完成再将转换为 string, 无论哪种转换, 都会重新分配内存, 并复制字节数组

11. 类型转换:
	只有强类型可以转换 没有隐式转换
强制类型转换的基本语言:
	T(表达式)

*/

func main() {
	s := "hello,world"
	fmt.Println(len(s))

	// 分割字符串
	fmt.Println(strings.Split(s, ",")[0])

	// 是否包含
	fmt.Println(strings.Contains(s, "hello"))

	if strings.Contains(s, "hello") {
		fmt.Println("包含")
	} else {
		fmt.Println("不包含")
	}

	// 前缀 后缀判断
	if strings.HasPrefix(s, "hello") {
		fmt.Println("前缀包含")
	}

	// 后缀判断
	if strings.HasSuffix(s, "world") {
		fmt.Println("后缀包含")
	}

	// 获取指定字符串第一次出现下标
	fmt.Println(strings.Index(s, "o"))

	// 获取最后一次出现的下标
	fmt.Println(strings.LastIndex(s, "o"))

	// 拼接字符串
	fmt.Println(strings.Join(strings.Split(s, ","), "-"))

	// 打印类型
	fmt.Printf("%T", s)

	// 字符类型
	//a := '你'
	//fmt.Printf("%T", a)

	// 遍历字符串

	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}

	//修改字符串
	//s1 := "白萝卜"
	//bytes1 := []byte(s1)
	//fmt.Println(bytes1)
	//bytes1[0] = 'H'
	//print(string(bytes1))
	//
	//s2 := "胡萝卜"
	//runes2 := []rune(s2)
	//runes2[0] = 'H'
	//print(string(runes2))

	// 类型转换

	var a, b = 10, 20
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("%T", c)

}
