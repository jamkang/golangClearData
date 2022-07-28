package network

type EventCallBack func()

type stEvent struct {
	event_type int

	//回调函数
	callback EventCallBack

	//Tcp消息
	cmd     int
	tcp_msg interface{}
	client  INet
}

type stMsgCallBack struct {
	cmd      int
	req_msg  interface{}
	callback interface{}
}

type stMainLoop struct {
	event_chan    chan *stEvent
	msg_callbacks map[int]*stMsgCallBack
}

//定义个全局变量存储注册的函数和cmd
var (
	MainLoop = &stMainLoop{
		event_chan:    make(chan *stEvent, 2000),
		msg_callbacks: make(map[int]*stMsgCallBack),
	}
)

//注册回调函数到msg_callbacks中
func (self *stMainLoop) RegistMsg(cmd int, callback interface{}, req_msg interface{}) {
	self.msg_callbacks[cmd] = &stMsgCallBack{cmd: cmd, callback: callback, req_msg: req_msg}
}
