package main

import (
	"errors"
	"net/rpc"
	"net"
	"fmt"
	"os"
)

//rpc参数：A除以B
type Args struct {
	A, B int
}

//rpc返回值：Quo为商，Rem为余数
type Quotient struct {
	Quo, Rem int
}

type Arith int

func (arith *Arith) Multiply(args Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (arith *Arith) Divide(args Args, quotient *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quotient.Quo = args.A / args.B
	quotient.Rem = args.A % args.B
	return nil
}

func main(){
	//在rpc模块上注册一个算法服务
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddress, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddress)
	checkError(err)

	for{
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//每个tcp连接对应一个goroutine，从而实现多用户并发处理
		go rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}