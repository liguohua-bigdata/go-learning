package main
import (
	"fmt"
)
func main() {
	//括号{}中的值的数量不能大于在方括号[]中为数组声明指定的元素数量。
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	//如果省略数组的大小，则只创建一个足够容纳初始化的数组
	var balance2 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance2)
}