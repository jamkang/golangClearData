package io_te

import (
	"fmt"
	"io"
	"os"
)

//封装一下io中的输入输出
type FooReader struct {
}

func (self *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in >")
	return os.Stdin.Read(b)
}

type FooWrite struct {
}

func (self *FooWrite) Write(b []byte) (int, error) {
	fmt.Print("out >")
	return os.Stdout.Write(b)
}

//对以上的调用

func UserFoo() {
	var (
		read  FooReader
		write FooWrite
	)
	input := make([]byte, 1024)
	s, err := read.Read(input)
	if err != nil {
		fmt.Println("in error:", err)
	}
	fmt.Printf("in byte nums is %d\n", s)
	s, err = write.Write(input)
	if err != nil {
		fmt.Println("out error:", err)
	}
	fmt.Printf("out byte nums is %d\n", s)
}

//io.copy 效果是上面调用的一样，但是回一直运行程序，不让他关闭
func UserCopyFoo() {
	var (
		read  FooReader
		write FooWrite
	)
	_, err := io.Copy(&write, &read)
	if err != nil {
		fmt.Println("io err:", err)
	}
}
