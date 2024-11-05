package main

import (
	"log"
)

func main() {
	// 调用 wire 生成的 InitializeApp 函数
	app, err := InitializeMyApp()
	if err != nil {
		log.Fatalf("应用程序初始化失败: %v", err)
	}

	// 启动服务器
	if err := app.Run(); err != nil {
		log.Fatalf("服务器运行失败: %v", err)
	}

}
