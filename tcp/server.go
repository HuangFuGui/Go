package main

import (
	"net"
	"time"
	"fmt"
	"strings"
	"strconv"
	"io"
	"os"
	"log"
)

const CLIENTCLOSED  = "客户端已关闭连接"

func main(){
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1234")
	checkError(err)

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	defer conn.Close()
	for {
		request := make([]byte, 128)//防止洪水攻击，并且每次要清空数组，因为conn.Read会把内容append到原来的切片
		read_len, err := conn.Read(request)
		if err != nil {
			if err == io.EOF {
				fmt.Println(CLIENTCLOSED)
			} else {
				fmt.Println(err)
				//输出i/o timeout，相当于i/o服务时间，客户端一直没来数据或来数据来得太慢了
				//一般我们浏览器网速太慢会得到i/o timeout的返回值
			}
			break
		}
		if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
	}
}

func checkError(err error){
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}