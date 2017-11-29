//defer、return、返回值之间执行顺序：https://my.oschina.net/henrylee2cn/blog/505535
package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("a return: ", a())
	fmt.Println("b return: ", b())

	c := c()
	fmt.Println("c return: ", *c, c)

	defer P(time.Now())
	time.Sleep(5e9)
	fmt.Println("main ", time.Now())
	/*输出：
	main  2017-10-12 14:22:12.1258399 +0800 CST
	defer 2017-10-12 14:22:07.1252229 +0800 CST
	P     2017-10-12 14:22:12.1688339 +0800 CST
	*/
}

//匿名返回值的情况
func a() int {
	var i int = 10
	defer func(){
		i++
		fmt.Println("a defer2: ", i)
	}()
	defer func(){
		i++
		fmt.Println("a defer1: ", i)
	}()
	return i * 2
	/*输出：
	a defer1:  11
	a defer2:  12
	a return:  20
	*/
}

//有名返回值的情况
func b() (i int) {
	i = 10
	defer func(){
		i++
		fmt.Println("b defer2: ", i)
	}()
	defer func(){
		i++
		fmt.Println("b defer1: ", i)
	}()
	return i * 2
	/*输出：
	b defer1:  21
	b defer2:  22
	b return:  22
	*/
}

func c() *int {
	var i int = 10
	defer func(){
		i++
		fmt.Println("c defer2: ", i, &i)
	}()
	defer func(){
		i++
		fmt.Println("c defer1: ", i, &i)
	}()
	return &i
	/*输出：
	c defer1:  11 0xc04203a280
	c defer2:  12 0xc04203a280
	c return:  12 0xc04203a280
	*/
}

func P(t time.Time) {
	fmt.Println("defer", t)
	fmt.Println("P    ", time.Now())
}