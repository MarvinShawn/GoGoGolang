package main

import "fmt"


//实现一个 fibonacci 函数，它返回一个函数（闭包）
// 该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
func fibonacci() func() int  {

	x,y := 0,1

	return func() int {
		sum := x
		x,y = y,x+y
		return sum
	}

}


func main() {

	f := fibonacci()

	for i := 0; i < 10 ; i++  {
		fmt.Println(f())
	}
	
}
