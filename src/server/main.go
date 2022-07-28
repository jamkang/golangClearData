package main

import (
	"bufio"
	"fmt"
	"github.com/golang/glog"
	"net"
)

func main() {
	listen, e := net.Listen("tcp", "127.0.0.1:20000")
	if e != nil {
		glog.Error("listen fail,err:", e)
	}
	for {
		conn, e := listen.Accept()
		if e != nil {
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, e := reader.Read(buf[:])
		if e != nil {
			fmt.Println("read from client failed err:", e)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client的数据", recvStr)
		conn.Write([]byte(recvStr)) //发送数据
	}
}
