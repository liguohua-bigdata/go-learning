package main

import "fmt"

func main() {
	sum := 0
	//for循环不要（），但循环体必须用{ }括起来。
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
