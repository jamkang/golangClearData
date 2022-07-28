package test

import (
	"encoding/json"
	"fmt"
)

/*

 */
//代表最里层的结构体
type appInfo struct {
	Appid string `json:"appId"`
}

//代表第二层的结构体
type response struct {
	RespCode string    `json:"respCode"`
	RespMsg  string    `json:"respMsg"`
	AppInfo  []appInfo `json:"app"`
}

//代表最外层花括号的结构体
type JsonResult struct {
	Resp response `json:"resp"`
}

type Sthandler struct {
	Name string `json:"name"`
	appInfo
}

//嵌套结构体数组，
//func UserDataStruct() {
//	jsonStr := `{"resp": {"respCode": "000000","respMsg": "成功","app": [{"appId": "d12ab8"},{"appId": "d12abd47e6a"},{"appId": "sdvfg30b8"}]}}`
//
//	var JsonRes JsonResult
//	if err := json.Unmarshal([]byte(jsonStr), &JsonRes); err != nil {
//		fmt.Println("fail", err)
//	} else {
//		fmt.Println("after parse", JsonRes.Resp.AppInfo[2])
//	}
//}

//嵌套一个继承的结构体
func UesrInhertStruct() {
	jsonstr := `{"name":"wangkang","appId":"12"}`

	var stHandler Sthandler
	if err := json.Unmarshal([]byte(jsonstr), &stHandler); err != nil {
		fmt.Println("fail", err)
	} else {
		fmt.Println("after parse", stHandler)
	}
}

//----------------------------------------------------------------------------------
