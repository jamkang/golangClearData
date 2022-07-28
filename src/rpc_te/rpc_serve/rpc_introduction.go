package rpc_serve

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"net"
	"project01/src/rpc_te/protobuf_te"
)

type HelloService struct {
}

//1.0.2所加
//func RegisterHelloSerVice(name string) error {
//	return rpc.RegisterName(name+"HelloService", new(HelloService))
//}

//func (p *HelloService) Hello(request string, reply *string) error {
//	*reply = "hello:" + request
//	return nil
//}

//利用protobuf的String重新实现HelloService
//func (p *HelloService) Hello(request *protobuf_te.String, reply *protobuf_te.String) error {
//	reply.Value = "hello:" + request.GetValue()
//	return nil
//}
//func (p *HelloService) InitServer(into string, replay *string) error {
//	*replay = into + "被初始化了"
//	return nil
//}

/////*
////1.0.1简单的使用rpc的客户端
////*/s
//func UserHello() {
//	rpc.RegisterName("HelloService", new(HelloService))
//	listener, err := net.Listen("tcp", ":1234")
//	if err != nil {
//		log.Fatal("ListenTCP error:", err)
//	}
//	conn, err := listener.Accept()
//	if err != nil {
//		log.Fatal("Accept error:", err)
//	}
//	rpc.ServeConn(conn)
//}

/*
StandardUser
@author: 王康
@time: 2022-07-06 18:55:01
@Description: 1.0.2规范的设计一个服务段
*/
//func StandardUse() {
//	RegisterHelloSerVice("test/")
//	listrner, err := net.Listen("tcp", ":1234")
//	if err != nil {
//		log.Fatal("ListenTCP error:", err)
//	}
//	for {
//		conn, err := listrner.Accept()
//		if err != nil {
//			log.Fatal("Accept error:", err)
//		}
//		go rpc.ServeConn(conn)
//	}
//
//}

/*
1.1.1简单的使用rpc的不同语言之间的使用 服务端，
使用了系统包net/rpc/jaonrpc
*/
//func UserHello() {
//	rpc.RegisterName("HelloService", new(HelloService))
//	listener, err := net.Listen("tcp", ":1234")
//	if err != nil {
//		log.Fatal("ListenTCP error:", err)
//	}
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			log.Fatal("Accept error:", err)
//		}
//		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
//		//不同语言之间调用用 rpc.ServeCodec 函数替代了
//		//rpc.ServeConn 函数，传入的参数是针对服务端的 json 编解码器。
//	}
//}

//-----------------------------------------------------------------------------------------------
//1.2.1 实现了protobuf中的接口，完成rpc
type ProdService struct {
}

/*
注意，在后面添加了 args ... 就是实现了client的接口
*/
func (self *ProdService) GetProdStock(ctx context.Context, in *protobuf_te.ProdRequest) (*protobuf_te.ProdResponse, error) {
	return &protobuf_te.ProdResponse{ProdStock: 20}, nil
}

func ProfuncServer() {
	rpcSer := grpc.NewServer()
	protobuf_te.RegisterProServiceServer(rpcSer, new(ProdService))

	//监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("监听有问题", err)
	}
	rpcSer.Serve(listen)
}

// -----------------------------------------------------------------------------------------------
//1.2.2 完成了一个TLS认证服务端，需要用到1.2.1中的内容
func GrpcTLSUse() {
	//监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("监听有问题", err)
	}
	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("../OpenSSl/server.pem", "../OpenSSl/server.key")
	if err != nil {
		fmt.Println("TLS认证出现问题", err)
	}
	// 实例化grpc Server, 并开启TLS认证
	rpcSer := grpc.NewServer(grpc.Creds(creds))

	protobuf_te.RegisterProServiceServer(rpcSer, new(ProdService))
	rpcSer.Serve(listen)
}

// -----------------------------------------------------------------------------------------------
//1.2.3 完成了一个TLS双向认证的认证服务端，需要用到1.2.1中的内容
func TwoWayTLSAuth() {
	//从证书相关文件中读取和解析信息,得到证书公钥,密钥对
	cert, err := tls.LoadX509KeyPair("../OpenSSl/server.pem", "../OpenSSl/server.key")
	if err != nil {
		fmt.Println("证书读取错误", err)
	}
	//创建一个新的,空的 Certpool
	certp := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../OpenSSl/ca.crt")
	if err != nil {
		fmt.Println("ca证书读取失败", err)
	}
	//尝试解析所传入的PEM编码的证书.如果解析成功会将其加到Certpool中,便于后面的使用
	certp.AppendCertsFromPEM(ca)
	//构建基于TLS的TransportCredentials选项
	creds := credentials.NewTLS(&tls.Config{
		//设置证书链,允许包含一个或者多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端的证书,可以根据实际情况选用以下的参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		//设置根证书的集合,验证方式使用ClientAuth中设定的模式
		ClientCAs: certp,
	})

	// 实例化grpc Server, 并开启TLS认证
	rpcSer := grpc.NewServer(grpc.Creds(creds))

	protobuf_te.RegisterProServiceServer(rpcSer, new(ProdService))
	//监听
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("监听有问题", err)
	}
	rpcSer.Serve(listen)
}

//-------------------------------------------------------------------------------------------------------------

// -----------------------------------------------------------------------------------------------
//1.2.4 完成了一个Token认证服务端，需要用到1.2.1中的内容
//基于1.2.2版本的基础上添加Token认证

func UserAuth(ctx context.Context) error {
	//ctx实际上拿到传输的用户名和密码
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("missing credentials")
	}
	var user, passwd string
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["passwd"]; ok {
		passwd = val[0]
	}
	if user != "admin" || passwd != "1234" {
		return fmt.Errorf("auth wrongful")
	}
	return nil
}

func GrpcTokenUse() {
	//从证书相关文件中读取和解析信息,得到证书公钥,密钥对
	cert, err := tls.LoadX509KeyPair("../OpenSSl/server.pem", "../OpenSSl/server.key")
	if err != nil {
		fmt.Println("证书读取错误", err)
	}
	//创建一个新的,空的 Certpool
	certp := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../OpenSSl/ca.crt")
	if err != nil {
		fmt.Println("ca证书读取失败", err)
	}
	//尝试解析所传入的PEM编码的证书.如果解析成功会将其加到Certpool中,便于后面的使用
	certp.AppendCertsFromPEM(ca)
	//构建基于TLS的TransportCredentials选项
	creds := credentials.NewTLS(&tls.Config{
		//设置证书链,允许包含一个或者多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端的证书,可以根据实际情况选用以下的参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		//设置根证书的集合,验证方式使用ClientAuth中设定的模式
		ClientCAs: certp,
	})

	//实现token认证,需要合法的用户名密码
	//实现一个拦截器
	var authInterRpc grpc.UnaryServerInterceptor
	authInterRpc = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = UserAuth(ctx)
		if err != nil {
			return
		}
		return handler(ctx, req)
	}
	// 实例化grpc Server, 并开启TLS认证 ,
	//使用grpc.UnaryInterceptor()把拦截器放入连接中
	rpcSer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(authInterRpc))

	protobuf_te.RegisterProServiceServer(rpcSer, new(ProdService))
	//监听
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("监听有问题", err)
	}
	rpcSer.Serve(listen)

}

//-------------------------------------------------------------------------------------------------------------
