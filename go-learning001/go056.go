package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int;

	for  i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}

	for  i = 0; i < MAX; i++ {
		//address
		fmt.Printf("addre of a[%d] = %d", i,ptr[i] )
		//value
		fmt.Printf("Value of a[%d] = %d", i,*ptr[i] )
		fmt.Println()
	}
}