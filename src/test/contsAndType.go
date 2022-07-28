package test

//类型重定义后是不是还是和原来相等

type currencyType int

const (
	RMB currencyType = iota
	DOLLAR
)

func LabelCurrent() {
	var te01 int = 1
	if currencyType(te01) == DOLLAR {
	}
}
