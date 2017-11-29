package main

import (
	"os"
	"fmt"
	"net/rpc"
	"log"
)

//rpc参数：A除以B
type Args struct {
	A, B int
}

//rpc返回值：Quo为商，Rem为余数
type Quotient struct {
	Quo, Rem int
}

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	//DialHTTP会在net.Dial之后向conn写入connect path http的字节流信息，
	//使得服务端可以将通过tcp收到的字节流向上抽象得到http请求对象，走http协议去处理请求
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quotient Quotient
	err = client.Call("Arith.Divide", args, &quotient)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quotient.Quo, quotient.Rem)
}