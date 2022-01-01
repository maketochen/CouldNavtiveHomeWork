package main

/*
编写一个 HTTP 服务器，
接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200
*/
func main() {
	app := App{}
	app.Initialize(getEnv())
	app.Run(":8000")
}
