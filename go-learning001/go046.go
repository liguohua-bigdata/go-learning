package main

import "fmt"

//符串文字是不可变的，因此一旦创建后，字符串文字就不能更改了。

func main() {
	var greeting = "Hello world!"

	fmt.Printf("normal string: ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}
	fmt.Printf("\n")

	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	/*q flag escapes unprintable characters, with + flag it escapses non-ascii characters as well
	  to make output unambigous  */
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")
}
