package main

import (
	"errors"
	"net/rpc"
	"net/http"
	"fmt"
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
	//将rpc服务注册到http服务上，这样可以复用http的接口，减少代码量
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}