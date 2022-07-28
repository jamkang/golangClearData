package network

import (
	"encoding/binary"
	"io"
	"time"
)

type (
	INet interface {
		SetProperty(interface{}, interface{}) //设置属性
		GetProperty(interface{}) interface{}  //获取属性
		RemoveProperty(interface{})           //移除链接属性
		GetConnectId() string                 //获取连接ID
		SendMessage(IMessage)                 //发送消息
		RemoteAddr() string                   //获取IP
		Close()                               //关闭连接
	}
	INetConnector interface {
		INet
		Reconnect(time.Duration) //重连
	}

	// 事件
	INetHook interface {
		OnClose(INet) int              //连接关闭
		OnBackClose(INetConnector) int //连接(主动方)被关闭
		OnBackConnected(INet) int      //连接(主动方)已连接
	}

	IParser interface {
		NewMessage() IMessage          //新建
		SetByteOrder(binary.ByteOrder) //设置字节序
	}

	IMessage interface {
		GetId() int                // 兼容旧cmd
		Id() string                // 获取协议名字
		Bytes() []byte             // 获取整个报文
		Body() []byte              // 获取包体
		ReadFrame(io.Reader) error // 读包
		Encrypt() bool             // 加密
		Decrypt() bool             // 解密
	}
)
