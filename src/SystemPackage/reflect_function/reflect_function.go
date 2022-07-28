package reflect_function

import (
	"fmt"
	"reflect"
)

var (
	funcda = make(map[int]interface{})
)

func reflect_function() {
	//函数可以当作一个值赋值为一个变量
	funcda[0] = ReginstMsg

	//查看类型
	fmt.Println("这个类型为：", reflect.TypeOf(funcda[0]))
	//对该函数进行调用
	revalue := reflect.ValueOf(funcda[0])
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(6)
	revalue.Call(in)
}

func ReginstMsg(data interface{}) {
	fmt.Println("this is a function ", data)
}
