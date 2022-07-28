package test

import (
	"fmt"
	"sort"
)

//关于结构体切片的排序

type User struct {
	index int //排序字段
	name  string
}

func SortStructSlice() {
	user := []User{
		{index: 1, name: "wang"},
		{index: 5, name: "kand"},
		{index: 2, name: "xigua"},
	}
	sort.Slice(user, func(i, j int) bool {
		return user[i].index < user[j].index
	})
	for _, v := range user {
		fmt.Println(v)
	}
}
