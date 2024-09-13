package main

// 测试做服务拆分
func main() {
	server := NewApp()
	server.Start()
}
