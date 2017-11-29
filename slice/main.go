package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	//slice重组
	//s := make([]int, 1, 1024)
	//s = s[:len(s)+1]
	//s[1] = 2
	//fmt.Println(s)

	//如果append goroutine在完成前不让出cpu，则线程安全，否则panic: runtime error: index out of range
	//s := make([]int, 0, 100)
	//go func() {
	//	for i := 0; i < 1024; i++{
	//		s = append(s, i)
	//	}
	//}()
	//go func() {
	//	i := 0
	//	for {
	//		v := s[i]
	//		fmt.Println(v)
	//		i++
	//		if i == 1024 {
	//			i = 0
	//		}
	//	}
	//}()
	//go func() {
	//	i := 0
	//	for {
	//		v := s[i]
	//		fmt.Println(v)
	//		i++
	//		if i == 1024 {
	//			i = 0
	//		}
	//	}
	//}()
	//go func() {
	//	i := 0
	//	for {
	//		v := s[i]
	//		fmt.Println(v)
	//		i++
	//		if i == 1024 {
	//			i = 0
	//		}
	//	}
	//}()
	//time.Sleep(5 * time.Second)
	//fmt.Println("finished")

	//slice丢失修改：当A和B两个协程运行append的时候同时发现s[1]这个位置是空的，
	//他们就都会把自己的值放在这个位置，这样他们两个的值就会覆盖，造成数据丢失
	//输出：954、976、963...
	//s := make([]int, 0, 1000)
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		//append其实是针对len的赋值操作，当底层数组长度cap不够了，才会重新申请分配连续空间
	//		s = append(s, i)
	//	}()
	//}
	//time.Sleep(5 * time.Second)
	//fmt.Println(len(s))

	//每次都正常输出1000
	s := make([]int, 0, 1000)
	var mux sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mux.Lock()
			s = append(s, i)
			mux.Unlock()
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(len(s))
}