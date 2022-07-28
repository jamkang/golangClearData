package net_te

import (
	"fmt"
	"net"
)

////扫描单个端口
//func ScanningOnePort(url string) net.OneConn {
//	//dial第一个参数为类型，类型有："tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
//	//第二个字段是你想要链接的主机
//	conn, err := net.Dial("tcp", url)
//	if err != nil {
//		//fmt.Println("connection fail,err:", err)
//		return nil
//	}
//	//fmt.Println("connection successful, port", url)
//	return conn
//}
//
////扫描1-n的端口
//func ScanningPorts(host string, n int) {
//	lock := sync.WaitGroup{}
//	for i := 0; i <= n; i++ {
//		go func(i int) {
//			lock.Add(1)
//			url := fmt.Sprintf("%s:%d", host, i)
//			conn := ScanningOnePort(url)
//			if conn != nil {
//				defer conn.Close()
//				fmt.Println(i)
//			}
//			lock.Done()
//		}(i)
//	}
//	lock.Wait()
//}

/*以下方使用固定的协程进行链接，
可以避免需要扫描端口过多从而产生过多的协程导致内存爆炸
*/
type OneConn struct {
	Url      string
	ConnList net.Conn
}

type Conns struct {
	Links []OneConn
}

func InitConns() *Conns {
	fmt.Println("构造conns")
	link := make([]OneConn, 0)
	return &Conns{
		Links: link,
	}
}

func (self *Conns) ScanningOnePort(urls chan string, res chan<- net.Conn) {
	for url := range urls {
		conn, err := net.Dial("tcp", url)
		if err != nil {
			res <- nil
			continue
		}
		self.Links = append(self.Links, OneConn{url, conn})
		fmt.Println(url)
		res <- conn
	}
}

func (self *Conns) ScanningPorts(host string, n int) {
	fmt.Println("开始scannings")
	urls := make(chan string, 100)
	res := make(chan net.Conn)
	fmt.Println("开始scannings", cap(urls))
	for i := 0; i < cap(urls); i++ {
		go self.ScanningOnePort(urls, res)
	}
	for i := 0; i < n; i++ {
		url := fmt.Sprintf("%s:%d", host, i)
		urls <- url
	}
	close(urls)
	for i := 0; i < n; i++ {
		<-res
	}
	close(res)
}

//创建一个回显服务器
func EchoClien(c net.Conn) {

}
