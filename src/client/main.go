package main

import (
	"bufio"
	"fmt"
	"github.com/golang/glog"
	"net"
	"os"
	"strings"
)

func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:20000")
	if e != nil {
		glog.Errorf("client connect fail,error_te:", e)
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, e = conn.Write([]byte(inputInfo))

		if e != nil {
			return
		}
		buf := [512]byte{}
		n, e := conn.Read(buf[:])
		if e != nil {
			fmt.Println("recv fail,err:", e)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
