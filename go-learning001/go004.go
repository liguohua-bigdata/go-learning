package main

import "fmt"

//函数多返回值
func swap(x int, y int) (int, int) {
	return y, x
}
func main() {
	a, b := swap(12, 13);
	fmt.Println(a, b)
}
