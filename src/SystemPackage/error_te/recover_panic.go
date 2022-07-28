package error_te

import (
	"errors"
	"fmt"
)

//当希望将捕获到的异常转为错误时，如果希望忠实返回原始的信息，需要针对不同的类型分别处理：
func recoverErr(err error) {
	if r := recover(); r != nil {
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = fmt.Errorf("Unknown error")
		}
	}
}
func foo() (err error) {
	defer recoverErr(err)
	panic("12")
}
