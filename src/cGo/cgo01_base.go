package cGo

import "C"

func base() {
	C.puts(C.CString("hello,world\n"))
}
