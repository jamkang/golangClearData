package rpc_client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"project01/src/rpc_te/protobuf_te"
)

// 1.0.2 start
//var _HelloServiceInterface = (*HelloServiceClient)(nil)
//
//const HelloServiceName = "path/to/pkg."
//
//type HelloServiceClient struct {
//	*rpc.Client
//}
//
//func (p *HelloServiceClient) Hello(request string, reply *string) error {
//	return p.Client.Call(HelloServiceName+"HelloService.Hello", request, reply)
//}

//1.0.2 end

/*
1.0.1简单的使用客户端
*/
//func SimpleteUse() {
//	client, err := rpc.Dial("tcp", "localhost:1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//	var reply string
//	err = client.Call("HelloService.InitServer", "abc", &reply)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("replay:", reply)
//}

/*
StandarUser
@author: 王康
@time: 2022-07-06 19:08:10
@Description: 1.0.2规范使用，防止定义名字重复和混乱使用
*/
//func DialHelloService(network, address string) (*HelloServiceClient, error) {
//	c, err := rpc.Dial(network, address)
//	if err != nil {
//		return nil, err
//	}
//	return &HelloServiceClient{Client: c}, nil
//}
//
//func StandarUse() {
//	client, err := DialHelloService("tcp", "localhost:1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//	var replay string
//	err = client.Hello("hello", &replay)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

/*
1.1.1简单的使用跨语言调用 客户端
先手工调用 net.Dial 函数建立 TCP 链接，然后基于该链接建立针对客户端的 json 编解码器。
*/
//func SimpleteUse() {
//	conn, err := net.Dial("tcp", "localhost:1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
//	var reply string
//	fmt.Println(client)
//	err = client.Call("HelloService.Hello", "abc", &reply)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(client)
//	fmt.Println("replay:", reply)
//}

//----------------------------------------------------------------------------------------------------
/*
1.2.1,实现protobuf实现一个rpc
*/

func ProClient() {
	// 不添加这个参数会报错，
	//grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		fmt.Println("链接错误", err)
	}
	defer conn.Close()

	prodClient := protobuf_te.NewProServiceClient(conn)
	rodRes, err := prodClient.GetProdStock(context.Background(), &protobuf_te.ProdRequest{
		ProdId: 12,
	})
	if err != nil {
		fmt.Println("调用失败", err)
	}
	fmt.Println(rodRes.ProdStock)
}

//-----------------------------------------------------------------------------------------
//1.2.2,实现TLS认证f实现一个grpc,关于SSL的公钥私钥在openssl中有
func GrpcUseClientTSL() {
	// TLS连接  第二个参数为服务器地址
	creds, err := credentials.NewClientTLSFromFile("../OpenSSl/server.pem", "server.com")
	if err != nil {
		fmt.Println("Failed to create TLS credentials ", err)
	}

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("链接错误", err)
	}

	defer conn.Close()

	//初始化客户端
	prodClient := protobuf_te.NewProServiceClient(conn)

	//然后就可以正常进行调用啦
	rodRes, err := prodClient.GetProdStock(context.Background(), &protobuf_te.ProdRequest{
		ProdId: 12,
	})
	if err != nil {
		fmt.Println("调用失败", err)
	}
	fmt.Println(rodRes)
}

//-----------------------------------------------------------------------------------------
//1.2.3,实现TLS双向认证实现一个grpc,关于SSL的公钥私钥在openssl中有\

func TwoAwayAuth() {
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
		ServerName: "*.server.com",
		RootCAs:    certp,
	})
	conn, err := grpc.Dial(":12345", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("链接错误", err)
	}

	defer conn.Close()
	//初始化客户端
	prodClient := protobuf_te.NewProServiceClient(conn)

	//然后就可以正常进行调用啦
	rodRes, err := prodClient.GetProdStock(context.Background(), &protobuf_te.ProdRequest{
		ProdId: 12,
	})
	if err != nil {
		fmt.Println("调用失败", err)
	}
	fmt.Println(rodRes)
}

//-----------------------------------------------------------------------------------------
//1.2.4,token认证,基于1,2,3的内容上进行添加Token功能

//客户端需要实现PerRPCCredentials接口
type UserInfo struct {
	User   string
	Passwd string
}

//相当于一个值传递
func (self *UserInfo) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"user": self.User, "passwd": self.Passwd}, nil
}

//RequireTransportSecurity指示凭据是否需要
////运输安全。
func (self *UserInfo) RequireTransportSecurity() bool {
	return false
}

func TokenAuth() {
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
		ServerName: "*.server.com",
		RootCAs:    certp,
	})
	//写如认证的参数
	user := &UserInfo{
		User:   "admin",
		Passwd: "1234",
	}
	//传入连接中
	conn, err := grpc.Dial(":12345", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(user))
	if err != nil {
		fmt.Println("链接错误", err)
	}

	defer conn.Close()
	//初始化客户端
	prodClient := protobuf_te.NewProServiceClient(conn)

	//然后就可以正常进行调用啦
	rodRes, err := prodClient.GetProdStock(context.Background(), &protobuf_te.ProdRequest{
		ProdId: 12,
	})
	if err != nil {
		fmt.Println("调用失败", err)
	}
	fmt.Println(rodRes)
}
