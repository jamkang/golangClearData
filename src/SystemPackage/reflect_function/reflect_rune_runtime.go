package reflect_function

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

//获取函数名字--------------------------------------------------------------------
func foo() {
}

//注意：rune 类型是 Go 语言的一种特殊数字类型，是类型 int32 的别名，用来区分字符值跟整数值
//可以用来截取带中文字符串，因为他刚好也占4字节，不会出现乱码,可以看做byte一样
//runtime包提供和go运行时环境的互操作，如控制go程的函数
func GetFunctionName(i interface{}, seps ...rune) string {
	// 获取完整的包名+函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	// 用 seps 进行分割 //和func Fields差不多
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

func GetFuncName() {
	// This will print "name: main.foo"
	fmt.Println("name:", GetFunctionName(foo))

	// runtime/debug.FreeOSMemory
	fmt.Println("FreeOSMemory01:", GetFunctionName(ReginstMsg))
	// FreeOSMemory
	fmt.Println("FreeOSMemory02:", GetFunctionName(ReginstMsg, '.'))
	// FreeOSMemory
	fmt.Println("FreeOSMemory03:", GetFunctionName(ReginstMsg, '/', '.'))
}

//------------------------------------------------------------------------------
