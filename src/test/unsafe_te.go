package test

import (
	"unsafe"
)

//失败，学习一下反射和unsafe

//func UseMenSet(s interface{}, c byte) {
//	svalue := reflect.ValueOf(s)
//	MemSet(unsafe.Pointer(&s), c, unsafe.Sizeof(s))
//	fmt.Printf("interface s:%T \n", svalue)
//	fmt.Println("type:", reflect.TypeOf(s))
//}

func MemSet(s unsafe.Pointer, c byte, n uintptr) {
	ptr := uintptr(s)
	var i uintptr
	for i = 0; i < n; i++ {
		pByte := (*byte)(unsafe.Pointer(ptr + i))
		*pByte = c
	}
}
