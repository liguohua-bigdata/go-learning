package main

import "fmt"

//函数定义
func add(x int, y int) int {
	return x + y
}

//函数定义，当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func add2(x, y int) int {
	return x + y
}
func main() {
	//函数调用
	fmt.Println(add(12, 34))

	//函数调用
	fmt.Println(add2(12, 34))
}
