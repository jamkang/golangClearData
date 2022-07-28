package io_te

import (
	"bufio"
	"log"
	"net"
)

func NetIo(conn net.Conn) {
	defer conn.Close()
	//conn实现了bufio中read和write的接口
	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("get client data fial , err:", err)
	}
	log.Println("write data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("write date to client fial ,err:", err)
	}
}
