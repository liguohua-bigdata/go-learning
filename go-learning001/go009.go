package main

/**
int，uint 和 uintptr 类型在32位的系统上一般是32位，而在64位系统上是64位。
当你需要使用一个整数类型时，应该首选 int，仅当有特别的理由才使用定长整数类型或者无符号整数类型。
 */
import (
	"fmt"
	"math/cmplx"
)
//同时与导入语句一样，变量的定义“打包”在一个语法块中。
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
