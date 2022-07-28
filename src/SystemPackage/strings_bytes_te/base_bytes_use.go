package strings_bytes_te

import (
	"bytes"
	"fmt"
)

/*
bytes该包比strings更加全面，读写为一体
*/

func BytesUser() {
	//利用字节切片或字符创建buffer
	buffer1 := bytes.NewBuffer([]byte("我是王康,"))
	buffer2 := bytes.NewBufferString("我在深圳")

	//写
	buffer1.Write([]byte("我12岁"))
	buffer2.WriteString("我在,读书")

	//读取内容
	ctx1 := make([]byte, 3)
	n, _ := buffer2.Read(ctx1)        //读取ctx大小的字节
	ctx2, _ := buffer1.ReadBytes(',') //遇到指定字节停止读取
	//buffer1.ReadString() 也一样，只不过返回字符串
	fmt.Println(string(ctx1))
	fmt.Println(string(ctx2))
	fmt.Println(n)
}
