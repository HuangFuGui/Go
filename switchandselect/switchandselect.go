package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)

	switch {
	//case <-ch: //  invalid case <-ch in switch (mismatched types int and bool)
	case <- ch == 1:
		fmt.Println("switch1")
		fallthrough//***无条件执行下一个case
	case <- ch == 2:
		fmt.Println("switch2")
	case <- ch==3:
		fmt.Println("switch3")
	}

	select {
	case v := <- ch:
		fmt.Println("select case1 v=", v)
	case v := <- ch:
		fmt.Println("select case2 v=", v)
	}
	/*输出:
	switch1
	switch2
	select case1 v= 2
	或：
	switch1
	switch2
	select case2 v= 2
	*/
}