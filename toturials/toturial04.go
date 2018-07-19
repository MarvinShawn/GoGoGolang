package main

import (
	"time"
	"fmt"
)

func spinner(delay time.Duration)  {

	for  {
		for _,r := range `-\|/` {
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}

}

func fib(x int) int  {

	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)


}


//func sum(a []int, c chan int) {
//		total := 0
//		for _,v := range a {
//			total += v
//		}
//
//		c <- total
//
//}i

func fibonaccii(n int,c chan int)  {

	x,y := 1,1
	for i:= 0; i < n ;i++{

		c <- x
		x,y = y,x+y

	}

	close(c)


}



func main() {


	c := make(chan int,10)
	go fibonaccii(cap(c),c)
	for i := range c {
		fmt.Println(i)
	}

	//a := []int{7,2,8,-9,4,0}
	//
	//c := make(chan int)
	//go sum(a[:len(a)/2],c)
	//go sum(a[len(a)/2:],c)
	//
	//x,y := <-c,<-c
	//
	//fmt.Println(x,y,x+y)

	//go spinner(10*time.Microsecond)
	//const n = 45
	//fibN := fib(n)
	//fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}
