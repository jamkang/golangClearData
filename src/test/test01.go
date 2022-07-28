package test

import "fmt"

func Test01(arg ...interface{}) {
	test02(arg...)
}

func test02(args ...interface{}) {
	for _, k := range args {
		fmt.Println(k)
	}
}
func Test03() {
	m := [...]int{ //数组语法的一种形式，数组长度由编译时推断出
		'b': 1,
		'c': 2,
		'd': 3,
	}
	m['a'] = 3
	m[0] = 12
	fmt.Println(len(m))
	for k, v := range m {
		fmt.Println(k, v)
	}
}

type Proson struct {
	name string
	Age  int
}

// go语言中私有和公有是按照包中进行的分类的，就算是结构体也是一样的
func (p *Proson) test02() {
	p.name = "wangk"
	p.Age = 23
}

func UserTest() {
}
