package os_te

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
文件操作：
1、打开或者创建
	Open(name string) (file *File, err error)
	OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	Create(name string) (file *File, err error)//创建文件
	Mkdir和 MkdirAll //创建文件夹
2、读：
	1、目录
		(f *File) Readdir(n int) (fi []FileInfo, err error)
		还可以使用到io/ioutil中的ReadDir
	2、文件
		(f *File) Read(b []byte) (n int, err error)
3、写：
	(f *File) Write(b []byte) (n int, err error) Write向文件中写入len(b)字节数据
	(f *File) WriteString(s string) (ret int, err error) WriteString类似Write，但接受一个字符串参数
	//追加时打开文件时 flag选择O_APPEND
4、文件路劲操作 path/filePath包中的 ，具体操作如下函数 UseFilePath
	对路径的一些处理
	文件路径查找、分割、拼整和路径下查找文件等
5、其他操作
	1、设置光标位置：func (f *File) Seek(offset int64, whence int) (ret int64, err error)
		Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：0为相对文件开头，
		1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。
	2、获取文件信息：func Stat(name string) (fi FileInfo, err error)
		FileInfo中含有文件信息，具体看ReadFileAndDir中介绍到
	3、判断文件目录是否存在，使用方法看如下函数 UserLabelExist
		IsNotExist(err error) bool
		IsExist(err error) bool
	4、文件重命名：func Rename(oldpath, newpath string) error //修改一个文件的名字，移动一个文件
	5、移除文件或文件夹：
		func Remove(name string) error //Remove删除name指定的文件或目录，目录中必须为空
		func RemoveAll(path string) error //RemoveAll删除path指定的文件，或目录及它包含的任何下级对象
*/

func OpenAndCreateFile(name string, flag int, perm os.FileMode) *os.File {
	file, _ := os.OpenFile(name, flag, perm)
	/*
		flag的选项
		O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
		O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
		O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
		O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
		O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
		O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
		O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
		O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
		可以选择多个，如os.O_CREATE|os.O_RDWR|os.O_APPEND

		perm的选项
		const (
		    // 单字符是被String方法用于格式化的属性缩写。
		    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
		    ModeAppend                                     // a: 只能写入，且只能写入到末尾
		    ModeExclusive                                  // l: 用于执行
		    ModeTemporary                                  // T: 临时文件（非备份文件）
		    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
		    ModeDevice                                     // D: 设备
		    ModeNamedPipe                                  // p: 命名管道（FIFO）
		    ModeSocket                                     // S: Unix域socket
		    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
		    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
		    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
		    ModeSticky                                     // t: 只有root/创建者能删除/移动文件
		    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
		    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
		    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
			相当于Linux中的文件权限。如0666等
		)
	*/

	file, _ = os.Open(name) //Open打开一个文件用于读取

	file, _ = os.Create(name) //Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件
	return file
}

func ReadFileAndDir(f *os.File, n int) {
	fileInfo, _ := f.Readdir(n) //返回一个有n个成员的[]FileInfo
	fmt.Println(fileInfo)
	/*
		type FileInfo interface {
		    Name() string       // 文件的名字（不含扩展名）
		    Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
		    Mode() FileMode     // 文件的模式位
		    ModTime() time.Time // 文件的修改时间
		    IsDir() bool        // 等价于Mode().IsDir()
		    Sys() interface{}   // 底层数据来源（可以返回nil）
		}
	*/
}

func UserLabelExist(name string) {
	//可以判断文件或者文件夹是否存在
	if _, err := os.Stat(name); os.IsNotExist(err) {
		fmt.Println(err)
	}
}

func UseFilePath(name string) {
	path, _ := filepath.Abs(".\\ostest.go")
	//Abs函数返回path代表的绝对路径，如果path不是绝对路径，会加入当前工作目录以使之成为绝对路径
	fmt.Println(path)
	//dir+base就是一个完整的路径
	fmt.Println(filepath.Base(path)) //Base函数返回路径的最后一个元素
	fmt.Println(filepath.Dir(path))  //Dir返回路径除去最后一个路径元素的部分

	fmt.Println("linux下：", filepath.Clean("..//.//sad//sad.png"))       //Clean函数去除多余的\和修正文件路劲
	fmt.Println("windows下：", filepath.FromSlash("..//.//sad//sad.png")) //??这个不知道怎么用？？？

	fmt.Println("文件后缀：", filepath.Ext(path)) //Ext函数返回path文件后缀
	//IsAbs(path string) bool IsAbs返回路径是否是一个绝对路径。

	fmt.Println("拼接路径:", filepath.Join("abc", "werwe", "dsaf.go")) //拼接路径
	//分割路径
	//最后一个路径分隔符后面位置分隔为两个部分（dir和file）并返回
	dir, file := filepath.Split(path)
	fmt.Println("分割路径1:", dir, "文件名：", file)
	//将PATH或GOPATH等环境变量里的多个路径分割开,以;分开
	dirlist := filepath.SplitList(path + ";" + path)
	fmt.Println("分割路径2:", dirlist)

	//Glob函数返回所有匹配模式匹配字符串pattern的文件或者nil（如果没有匹配的文件）
	//func Glob(pattern string) (matches []string, err error)
	fmt.Println("----------------Glob-------------")
	fmt.Println(filepath.Glob("./*.go"))
	fmt.Println(filepath.Glob("./os_test*"))
	fmt.Println("----------------Match------------")
	//查看格式是否匹配的
	fmt.Println(filepath.Match("./test/a.go", "./test/a.go"))
	fmt.Println(filepath.Match("./test/*.go", "./test/a.go"))
	fmt.Println(filepath.Match("./test/b.go", "./test/a.*"))
}
