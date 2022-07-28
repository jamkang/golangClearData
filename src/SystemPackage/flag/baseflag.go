package flag

import "flag"

type User struct {
	Name string
	Age  int
}

/*
flag包可以让参数传到变量中去
一般流程：
	1、函数注册flag，将变量与命令参数绑定
		有两种方式：
		1、DataType(name string, value bool, usage string)*datatype
		用指定的名称、默认值、使用信息注册一个type类型flag。返回一个保存了该flag的值的指针
		2、DataTypeVar(p *datatype, name string, value bool, usage string)
		指定的名称、默认值、使用信息注册一个type类型flag，并将flag的值保存到p指向的变量
	2、使用flag.parse从arguments中解析注册的flag
	3、就可以使用了
*/
func ArgsInitUser() *User {
	user := new(User)
	defer flag.Parse()
	flag.StringVar(&user.Name, "name", "未设置", "用户名")
	flag.IntVar(&user.Age, "age", 0, "用户年龄")
	return user
}
