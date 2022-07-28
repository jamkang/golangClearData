package strings_bytes_te

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
string好处就是申请一个内存反复使用，防止多次申请和使用内存
1、builder在内存中写
2、在内存中读取
*/

//生产一个n位的随机种子
func RandString(n int) string {
	chars := "qwertyuiopasdfghjklzxcvbnm1234567890"
	var builder strings.Builder
	rand.Seed(time.Now().UnixNano()) //添加种子，防止每一次随机数一致
	for i := 0; i < n; i++ {
		builder.WriteByte(chars[rand.Intn(len(chars))]) //writeByte 写一个字符到Builder中
	}
	return builder.String()
}

func Builder() {
	//定义strings.Builder 结构体对象,类似在内存中的流对象
	var builder strings.Builder
	builder.Write([]byte("我是wk\n")) //写如一个字符切片追加到 Builder 中
	builder.WriteString("是王康")      //写一个string到 Builder 中
	fmt.Println(builder.String())
	fmt.Println("len: ", builder.Len()) //计算 Builder 中的长度

	builder.Reset() //清空Builder
	fmt.Println("len: ", builder.Len())

	fmt.Println("随机种子为：", RandString(10))

}

func Reader(data string) {
	reader := strings.NewReader(data) //初始化字符串，创建Reader结构体指针

	var err error
	var n int
	for err == nil {
		ctx := make([]byte, 4)
		n, err = reader.Read(ctx) //读取len(ctx)长度的字节
		fmt.Println(string(ctx[:n]))
	}
}
