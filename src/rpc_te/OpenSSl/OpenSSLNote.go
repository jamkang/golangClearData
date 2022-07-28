package OpenSSl

/*
目前还是出现一些问题

 OpenSSL官网
官方下载地址： https://www.openssl.org/source/

OpenSSL官网没有提供windows版本的安装包，可以选择其他开源平台提供的工具。例如

http://slproweb.com/products/Win32OpenSSL.html

方法一：
生产OpenSSL步骤：
	1、进入bin目录，执行openssl
	2、执行openssl genrsa -des3 -out ca.key 2048(会生成server.key私钥文件)
	3、创建证书请求，openssl req -new -key ca.key -out ca.csr (ca.csr)
		其中common name 也就是域名：
	//4、删除密码 openssl rsa -in server.key -out server_no_passwd.key
	5、执行openssl x509 -req -days 365 -in ca.csr -signkey ca.key -out ca.crt
		（会生成server.crt）
自签证书完成

把配置文件复制出来,最后需要对openssl进行操作
打开req_extensions = v3_req # The extensions to add to a certificate request
打开copy_extensions = copy
在[ v3_req ]下添加subjectAltName = @alt_name
[ alt_name ]下添加DNS.1 = *.mszlu.com

	6. openssl genpkey -algorithm RSA -out server.key

	7. openssl req -new -nodes -key server.key -out server.csr -days 3650 -config ./openssl.cnf -extensions v3_req

	8. openssl x509 -req -days 365 -in server.csr -out server.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req

//---------------------------------------------------------------------------------------------------------------
2022、07、14的问题，留给以后解决
SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。
使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。

$ openssl req -new -sha256 \
    -key server.key \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=UnitedStack/OU=Devops/CN=localhost" \
    -reqexts SAN \
    -config <(cat /etc/pki/tls/openssl.cnf \
        <(printf "[SAN]\nsubjectAltName=DNS:localhost")) \
    -out ustack.csr
//----------------------------------------------------------------------------------------------------------------
*/
