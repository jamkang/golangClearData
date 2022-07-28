package TextTrans

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
)

/*
golang自带的json解析库encoding/json提供了json字符串到json对象的相互转换，
在json字符串比较简单的情况下还是挺好用的，但是当json字符串比较复杂或者嵌套比较多的时候，
就显得力不从心了，不可能用encoding/json那种为每个嵌套字段定义一个struct类型的方式，
这时候使用simplejson库能够很方便的解析。例如，有这样一个嵌套很深的json字符串，例子如下
*/
var json_str string = `{"rc" : 0,
  "error_te" : "Success",
  "type" : "stats",
  "progress" : 100,
  "job_status" : "COMPLETED",
  "result" : {
    "total_hits" : 803254,
    "starttime" : 1528434707000,
    "endtime" : 1528434767000,
    "fields" : [ ],
    "timeline" : {
      "interval" : 1000,
      "start_ts" : 1528434707000,
      "end_ts" : 1528434767000,
      "rows" : [ {
        "start_ts" : 1528434707000,
        "end_ts" : 1528434708000,
        "number" : "x12887"
      }, {
        "start_ts" : 1528434720000,
        "end_ts" : 1528434721000,
        "number" : "x13028"
      }, {
        "start_ts" : 1528434721000,
        "end_ts" : 1528434722000,
        "number" : "x12975"
      }, {
        "start_ts" : 1528434722000,
        "end_ts" : 1528434723000,
        "number" : "x12879"
      }, {
        "start_ts" : 1528434723000,
        "end_ts" : 1528434724000,
        "number" : "x13989"
      } ],
      "total" : 803254
    },
      "total" : 8
  }
}`

func simplateJson_te() {
	res, err := simplejson.NewJson([]byte(json_str))
	if err != nil {
		fmt.Println("err:", err)
	}
	rows, err := res.Get("result").Get("timeline").Array()
	fmt.Println(rows)
}
