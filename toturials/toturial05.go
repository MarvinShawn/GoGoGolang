package main

import (
	"fmt"
	"time"
	"os"
	"sync"
)

func Sum(s []int, c chan int)  {

	sum := 0
	for _,x := range s {
		sum += x
	}
	c <- sum

}

func main() {
	//ch := make(chan int)
	//
	////无缓存channel
	//slice := []int{1,2,3,4,5,6,7,8,9}
	//
	//go Sum(slice[0:len(slice)/2],ch)
	//go Sum(slice[len(slice)/2:],ch)
	//
	//x , y := <- ch , <-ch
	//println(x,y)



	//串联的channel
	// num := make(chan int)
	// sqr := make(chan int)

	 /*
	 go func() {
	 	for x :=0 ; x < 10 ;x++ {
	 		num <- x
	 		time.Sleep(time.Second)
		 }
		 close(num)
	 }()


	 go func() {
	 	for x := range num {
	 		sqr <- x * x
		}
		close(sqr)
	 }()

	for x:= range sqr {
		println(x)
	}
	 */
	 //go counter(num)
	 //go squarer(sqr,num)
	 //printer(sqr)
	 //
	 //loopgo()
	 //time.Sleep(time.Second)

	 //selectDemo()

	 //asyncExitDemo()

	//syncWaitDemo()
	//synxMutexDemo()
	//rwDemo()
	onceDemo()

}



// chan<- int 表示只写不读
func counter(out chan <- int)  {

	for x := 0;x < 10 ;x++ {
		out <- x
		time.Sleep(time.Second)
	}
	close(out)

}

// <-chan int 只读不写
func squarer(out chan<- int, in <-chan int)  {


	for x := range in {
		out <- x * x
	}
	close(out)

}

func printer(in <-chan int)  {

	for x := range in {
		fmt.Println(x)
	}

}


//并发循环
func loopgo()  {

	for i := 1;i < 10 ;i++  {
		go func(x int) {
			fmt.Printf("这是打印的第%d在执行! \n",x)
		}(i)
	}


}


// select 多路复用
func selectDemo()  {

	fmt.Println("程序开始")
	timeout := make(chan bool)
	ch := make(chan int)

	go func() {
		 time.Sleep(time.Second * 3)
		 timeout <- true
	}()

	for {
		select {
			case <-ch:
				fmt.Println("从ch中接收数据")
			case <-timeout:
				fmt.Println("超时时间到")
				return
		}
	}

}

// select 并发退出
func asyncExitDemo()  {

	fmt.Println("程序开始执行")

	exit := make(chan struct{})

	for i := 0; i <= 5 ;i++  {

		go func(x int) {

			for  {
				select {
					case <-exit:
						fmt.Println("第",x,"个goroutine退出")
						return
					default:
						fmt.Println("第",x,"个goroutine正在运行")
				}
				time.Sleep(time.Second)
			}
		}(i)

	}

	os.Stdin.Read(make([]byte,1))
	close(exit)
	time.Sleep(time.Second) //这句是为了让出时间片  不然 上面的goroutine 还没开始执行 main就跑完了
	fmt.Println("程序退出")


}

func syncWaitDemo()  {

	var wg sync.WaitGroup

	for i := 1;i < 6 ;i++  {
		wg.Add(1)
		go func(x int) {
			time.Sleep(time.Second)
			fmt.Println("第",x,"个goroutine执行完成")
			wg.Done()  // delta -1
		}(i)

	}
	wg.Wait()  // 阻塞直到 wg的delta为0
	fmt.Println("程序执行完成")

}

var sum int = 0
//互斥锁
func synxMutexDemo()  {

	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(2)
	go func() {
		for i:=0;i<1e8 ;i++  {
			m.Lock()
			sum++
			m.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i:=0;i<1e8 ;i++  {
			m.Lock()
			sum++
			m.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(sum)

}


//读写锁 写之间  读写之是互斥的  读之间不解锁
var rw sync.RWMutex

func writeData(wg *sync.WaitGroup,id int)  {
	rw.Lock()
	fmt.Println("第",id,"goroutine写加锁")
	for i := 0;i < 5 ;i++  {
		fmt.Print(id,"-w\n")
		time.Sleep(time.Second)
	}
	fmt.Println("\n第",id,"goroutine写解锁")
	rw.Unlock()
	wg.Done()

}

func readData(wg *sync.WaitGroup,id int)  {
	rw.RLock()
	fmt.Println("第",id,"goroutine读加锁")
	for i := 0;i < 5 ;i++  {
		fmt.Print(id,"-r\n")
		time.Sleep(time.Second)
	}
	fmt.Println("\n第",id,"goroutine读解锁")
	rw.RUnlock()
	wg.Done()

}

func rwDemo() {

	var wg sync.WaitGroup
	for i:=1;i < 3;i++{
		wg.Add(1)
		go writeData(&wg,i)
	}

	for i := 1;i < 6 ;i++  {
		wg.Add(1)
		go readData(&wg,i)
	}

	wg.Wait()
	fmt.Println("程序退出!")

}


func onceBody()  {
	fmt.Println("Only Once")
}

func onceDemo()  {
	var once sync.Once
	var wg sync.WaitGroup
	for i :=0;i < 10 ;i++  {
		wg.Add(1)
		go func() {
			once.Do(onceBody)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("程序结束")
}