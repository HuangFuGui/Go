//一个goroutine的异常没有被正确地恢复会造成整个程序崩溃：Process finished with exit code 2
//异常得到恢复后，程序能正常执行：Process finished with exit code 0
package main

import (
	"fmt"
	"errors"
	"time"
	"sync"
	"runtime/debug"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func(){
		err := funcA()
		if err == nil {
			fmt.Printf("err is nil\n")
		} else {
			fmt.Printf("err is %v\n", err)
		}
		wg.Done()
	}()
	go calculate()
	wg.Wait()
}

//err is foo，将异常信息显式的传递给错误
func funcA() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v\n", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			debug.PrintStack()
		}
	}()
	return funcB()
}

//err is nil，panic异常处理机制不会自动将异常信息传递给错误
//func funcA() error {
//	defer func() {
//		//recover只处理异常，不处理错误
//		if p := recover(); p != nil {
//			fmt.Printf("panic recover! p: %v\n", p)
//			debug.PrintStack()
//		}
//	}()
//	return funcB()
//}

func funcB() error {
	panic("foo")
	return errors.New("success")
}

func calculate(){
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("继续")
	}
	wg.Done()
}