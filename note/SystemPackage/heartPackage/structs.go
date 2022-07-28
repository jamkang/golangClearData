package heartPackage

import (
	"container/list"
	"project01/note/SystemPackage/xutil"
	"time"
)

type Server struct {
	Host       string
	Port       int
	WsPort     int
	Index      int
	GameId     int
	HostIndex  int32
	ServerType int32
	ServerId   int32
	Level      int
	GameType   int32
	Version    string
	// 配置文件
	ConfigFile string

	// 启用websocket服务
	Ws bool
	// 启用http服务
	Http bool
	// 启用rpc服务
	Rpc bool

	// *Config配置
	//cfg atomic.Value
	// xml通用服务配置

	// 处理器
	//Handler *Handler

	// 注册服务管理
	//RegisterMgr *RegisterMgr
	// 启动时间
	RunTime time.Time

	players map[int]*Player

	offlineSysMsg map[int]*list.List

	moneyFlag int
	boardFlag int

	//to do
	serverUpdating       bool
	whiteList            []int //白名单id，维护时间可进游戏测试
	mEarnLimit           int
	mBaseMoney           float64
	killplayer           map[int]float64
	inBlackRoomPlayer    map[int]float64
	mNeedKillPlayerCount int
	curMinPoint          float64 //当前被杀玩家中仇恨点数最低点数
	shareScore           float64
}

func NewServer(commandArgs *xutil.ServerCommandArgs) {

}

type Player struct {
}
