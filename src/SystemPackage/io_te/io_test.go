package io_te

import (
	"fmt"
	"testing"
)

func TestIo(t *testing.T) {
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
