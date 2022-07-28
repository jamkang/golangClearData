package test

import (
	"fmt"
	"testing"
)

type Test0 struct {
	Name string
	Age  int
}

func TestRand(t *testing.T) {
	data := make(map[string]string, 2)
	data["水果"] = "苹果"
	data["吸管"] = "不可"

	MapTest(data)
	for k, v := range data {
		fmt.Println(k, v)
	}
}

func MapTest(data map[string]string) {
	data["特i"] = "奥特曼"
}
