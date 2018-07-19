package main

import (
	"fmt"
	"runtime"
	"os"
)

//初始化逻辑包装为一个匿名函数处理
var pc [256]byte = func()(pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	return

}()


//这样的init初始化函数除了不能被调用或引用外，其他行为和普通函数类似。
// 在每个文件中的init初始化函数，在程序开始执行时按照它们声明的顺序被自动调用。
func init() {

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}


func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func f(x int) {

	fmt.Printf("f(%d)\n",x+0/x)
	defer fmt.Printf("defer %d\n",x)
	f(x-1)

}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

//这个例子中  返回值变量名为i，在defer表达式中可以修改这个变量的值。
// 所以，虽然在return的时候给返回值赋值为1，后来defer修改了这个值，
// 让i自增了1，所以，函数的返回值是2而不是1
func c() (i int) {
	defer func() { i++ }()
	return i
}


func d() {

	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("f")

}

//defer 语法
func main() {

	defer func() {
		fmt.Println("c")
		if err := recover(); err!=nil {
			fmt.Println(err)

		}
		fmt.Println("d")
	}()

	d()
}
