package json_te

import (
	"fmt"
)

func ReadCon() {
	fmt.Println("开始ReadCon")
	var data interface{}
	err := ReadJsonCfg("./conf/adFilter.conf", &data)
	if err != nil {
		fmt.Errorf("read ./conf/wordfilter.txt error")
		return
	}
	weis := data.(map[string]interface{})
	fmt.Println(weis["weis"])
}
