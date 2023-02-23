package main

/*
1. 值类型:
bool
int(32 or  64) int8 int16 int32 int64
uint(32 or 64) uint8 uint16 uint32 uint64
float32 float64
string
complex64 complex128
array  固定长度的数组


2. 引用类型
slice   序列数组
map 	映射
chan 	管道

3. 内置函数
append   用来追加元素到数组,slice 中 返回修改后的数组 slice
close    用来关闭管道
delete   从 map 中删除 key 对应的 value
panic	 停止常规的 goroutine (panic 和 recover  用来做错误处理)
recover  允许程序定义 goroutine 和 panic 的动作
imag    返回复数的虚部
real    返回复数的实部
make    用来分配内存,主要用来分配 slice, map, chan
new     用来分配内存,主要用来分配值类型,比如 int, struct
cap     用来返回 slice, map 最大长度
len     用来返回 slice, map, string, array 长度
print println 底层打印函数, 在部署环境中建议使用 fmt 包



内置接口 error
type error interface {
Error() string
}


*/
