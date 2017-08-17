package main

import "fmt"

//返回值可以被命名,返回值的名称应当具有一定的意义，可以作为文档使用。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	//没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
}

func main() {
	fmt.Println(split(17))
}