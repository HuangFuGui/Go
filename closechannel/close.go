//可以不断向已经关闭了的channel取数据（换个角度讲：此时channel已经“就绪”，在select中的case是可以执行的），只是这个数据为零值，且ok为false
package main

import "fmt"

func main(){
	ch := make(chan int)
	close(ch)//若注释掉close(ch)，协程会一直等待channel的数据，容易造成死锁，所以应该注意到任何时候从channel取数据，都有死锁的可能
	v, ok := <- ch
	fmt.Printf("%v %v\n", v, ok)
	/*输出：
	0 false
	*/
}