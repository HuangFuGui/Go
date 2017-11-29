// https://www.kancloud.cn/kancloud/effective/72199

package main

import (
	"fmt"
	"net/http"
)



type ByteSize float64

const (
	_           = iota 				// 0， ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)  // 1 << (10 * 1)
	MB								// 1 << (10 * 2)
	GB
	TB
	PB
	EB
	ZB
	YB
)

// 带缓冲区的channel可以像信号量一样使用，用来完成诸如吞吐率限制、并发控制等功能。
const MaxOutstanding = 1000

var sem = make(chan int, MaxOutstanding)



// 若switch后面没有表达式，它将匹配true，因此，我们可以将if-else-if-else链写成一个switch，这也更符合Go的风格。
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// switch case可通过逗号分隔来列举相同的处理条件。
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

// 跳出循环标签
func breakLoop() {
	Loop:
	for i := 1; i < 1000; i++ {
		switch {
		case i % 13 == 0:
			fmt.Println("求余13为0", i)
			break// 这个break其实意义不大
		case i % 33 == 0:
			fmt.Println("求余33为0", i)
			break Loop
		case i >= 0:
			fmt.Println("i >= 0", i)
		}
	}
	fmt.Println("跳出Loop")
}

// switch可用于判断接口变量的动态类型，如类型选择通过圆括号中的关键字type使用类型断言语法
func switchType(t interface{}) {
	switch t := t.(type) {
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	}

	/*
	var i interface{}
	i = 5
	v, ok := i.(int)
	fmt.Println(v, ok)// 5 true

	v, ok := i.(string)//  false
	*/
}



/*
entering: b
...
...
...
in b
entering: a
...
...
...
in a
leaving: a
leaving: b
*/
func trace(s string) string {
	fmt.Println("entering:", s)
	for i := 0; i < 3; i++ {
		fmt.Println("...")
	}
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}



func appendSliceForSlice() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
}



func process(r *http.Request) {

}

func ServeWithBug(queue chan *http.Request) {
	for req := range queue {
		<- sem
		go func() {
			process(req) // 有bug，Go中for循环的实现，循环的迭代变量会在循环中被重用，因此req变量会在所有Goroutine间共享。
			sem <- 1
		}()
	}
}

func Serve(queue chan *http.Request) {
	for req := range queue {
		<- sem
		go func(req *http.Request) { // 保证req变量是每个Goroutine私有的
			process(req)
			sem <- 1
		}(req)
	}
}

func handle(queue chan *http.Request) {
	for r := range queue {
		process(r)
	}
}

// 启动固定数量的handle Goroutine，每个Goroutine都直接从channel中读取请求。这个固定的数值就是同时执行process的最大并发数。
func ServeFixedHandler(clientRequests chan *http.Request, quit chan bool) {
	// Start handlers
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<- quit  // Wait to be told to exit.
}



func main(){
	// 反转数组a，Go没有逗号操作符，想要在for中使用多个变量，应采用平行赋值的方式
	a := [5]int{1, 2, 3, 4, 5}
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}