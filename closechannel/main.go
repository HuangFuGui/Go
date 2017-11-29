package main

import (
	"fmt"
	"time"
)

const SLEEPDURATION = time.Second * 2

func main(){
	ch := make(chan string)
	go sendData(ch)
	//getDataByOK(ch)
	getDataByRange(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	time.Sleep(SLEEPDURATION)
	ch <- "Tripoli"
	time.Sleep(SLEEPDURATION)
	ch <- "London"
	time.Sleep(SLEEPDURATION)
	ch <- "Beijing"
	time.Sleep(SLEEPDURATION)
	ch <- "Tokio"

	close(ch)//若注释掉close(ch)，两种情况都会deadlock
}

func getDataByOK(ch chan string) {
	for {
		v, ok := <- ch//只有close(ch)时才会有ok为false，否则将一直等待数据，容易造成死锁
		if !ok {
			break
		}
		//处理数据
		fmt.Printf("%s\n", v)
	}
}

func getDataByRange(ch chan string) {
	for v := range ch {//只有close(ch)时才会跳出循环，否则将一直等待数据，容易造成死锁
		//处理数据
		fmt.Printf("%s\n", v)
	}
}
