package heartPackage

import (
	"github.com/golang/glog"
	"project01/note/SystemPackage/heartPackage/network"
	"reflect"
)

func (self *Server) registMsgInTable(cmd int, callback interface{}, req_msg interface{}) {

}

//注册函数
func (self *Server) registMsgWithPlayer(cmd int, callback interface{}, regmsg interface{}) {
	callBackWrapper := func(client network.INet, reg_msg interface{}) {
		player := self.getPlayerByFd(client)
		if player == nil {
			glog.Errorf("player is nil")
			return
		}
		cd_value := reflect.ValueOf(callback)

		params := make([]reflect.Value, 2)
		params[0] = reflect.ValueOf(player)
		params[1] = reflect.ValueOf(regmsg)
		cd_value.Call(params)
	}
	network.MainLoop.RegistMsg(cmd, callBackWrapper, regmsg)

}

func (self *Server) getPlayerByFd(client network.INet) *Player {
	if client.GetProperty("uid") == nil {
		glog.Errorf("not logined yet")
		return nil
	}

	uid := client.GetProperty("uid").(int)
	if uid <= 0 {
		glog.Errorf("uid error,uid:%d", uid)
		return nil
	}

	return self.GetPlayer(uid)
}
