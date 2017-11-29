package main

import (
	"net"
	"fmt"
	"os"
)

func main(){
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1234")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("timestamp"))
	checkError(err)
	result := make([]byte, 128)
	read_len, err := conn.Read(result)
	checkError(err)
	fmt.Println(string(result[:read_len]))

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result = make([]byte, 128)
	read_len, err = conn.Read(result)
	checkError(err)
	fmt.Println(string(result[:read_len]))

	_, err = conn.Write([]byte("timestamp"))
	checkError(err)
	result = make([]byte, 128)
	read_len, err = conn.Read(result)
	checkError(err)
	fmt.Println(string(result[:read_len]))

	conn.Close()
	os.Exit(0)//程序正常退出

	//1.注意上述的conn.Read方法不能换成ioutil.ReadAll(conn)
	//ReadAll会等到error或EOF的时候才会返回，所有只有服务端主动
	//关闭连接客户端才会返回，否则会一直阻塞，得不到想要的效果
	//2.如果服务端发送的话，客户端要接收，否则程序会一直阻塞
}
