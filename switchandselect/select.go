//1. 从不同的并发执行的协程中获取值可以通过关键字select来完成，它的行为像是“你准备好了吗”的轮询机制；select监听进入通道的数据。
//2. 在select中，fallthrough 是不被允许的
//3. select 做的就是：选择处理列出的多个通信情况中的一个。
//	如果都阻塞了，会等待直到其中一个可以处理
//	如果多个可以处理，随机选择一个
//	如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。

package main

import (
	"time"
	"fmt"
)

const SLEEPDURATION = time.Second * 1

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(SLEEPDURATION)
}

func pump1(ch chan int) {
	for i := 0;; i++ {
		ch <- i * 2
	}
	//ch <- -1
	//close(ch)
}

func pump2(ch chan int) {
	for i := 0;; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for{
		select{
		case v, ok := <- ch1:
			if ok == false {
				fmt.Println("channel已经关闭")
				//return//会返回整个函数，这跟switch不一样
				continue//后面的代码不再执行，但继续for select循环
			}
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <- ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		//default://注意：default会跟就绪的case一起随机输出
		//	fmt.Printf("default\n")
		}
	}
}