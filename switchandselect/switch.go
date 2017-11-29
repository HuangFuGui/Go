//1. 一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 break 语句来表示结束。
//2. 如果在执行完每个分支的代码后，还希望继续执行后续分支的代码（***无条件执行），可以使用 fallthrough 关键字来达到目的。
//3. 您同样可以使用 return 语句来提前结束代码块的执行。当您在 switch 语句块中使用 return 语句，并且您的函数是有返回值的，
//   您还需要在 switch 之后添加相应的 return 语句以确保函数始终会返回。
package main

import "fmt"

func main(){
	var num1 int = 98
	switch num1 {
	case 97, 98, 99:
		fmt.Printf("It's equal to 97 or 98 or 99\n")
	case 100:
		fmt.Printf("It's equal to 100\n")
	default:
		fmt.Printf("default\n")
	}
}
