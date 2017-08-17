package main

import "fmt"
/**
在大多数操作系统上，程序不允许访问地址0处的内存，因为该内存是由操作系统保留的。
然而，存储器地址0具有特殊意义; 它表示指针不打算指向可访问的存储器位置。但是按照惯例，
如果指针包含nil(零)值，则假设它不指向任何东西。
 */
func main() {
	var ptr *int

	fmt.Printf("The value of ptr is : %x\n", ptr)
}
