package json_te

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJsonCfg(path string, data interface{}) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("can't load config file. path:%s,err:%s", path, err)
		return err
	}
	//dir, _ := os.Getwd()
	//fmt.Println(1, string(bs))
	err = json.Unmarshal(bs, data)
	if err != nil {
		fmt.Printf("parse config file error. path:%s,err:%s", path, err)
		return err
	}
	//fmt.Println(data)
	return nil
}
