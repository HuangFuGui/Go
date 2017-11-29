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
		fmt.Println("Usage:", os.Args[0], "server:port")
		os.Exit(1)
	}
	server := os.Args[1]

	client, err := rpc.Dial("tcp", server)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//Synchronous call
	args := Args{20, 8}
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
