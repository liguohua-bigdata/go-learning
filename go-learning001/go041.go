package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hippo", "Curry")
	fmt.Println(a, b)
}
