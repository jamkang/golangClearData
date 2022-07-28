package xutil

type IServer interface {
	OnRun()     // 运行服务
	OnDestroy() // 服务销毁
}

type ServerCommandArgs struct {
	// 平台
	GameId int
	// 版本号
	Version string
	// 服务类型
	GameType int
	// 绑定IP
	Host string
	// 绑定tcp端口
	Port int
	// ws端口(弃用)
	WsPort int
	// 是否开启websocket
	Ws bool
	// 是否开启rpc
	Rpc bool
	// 是否开启http
	Http bool
	// 服务索引
	Index int
	// level 区分玩法场次
	Level int
	// 配置文件路径
	ConfigFile string
	// http绑定地址
	HttpAddr string
	// rpc绑定地址
	RpcAddr string
	// server 环境 local:本地环境  dev:开发测试环境 test:测试环境  product：线上正式
	Env string

	JsonFile string

	ServerId int
}

var (
	CommandArgs = &ServerCommandArgs{}
)
