package main

import (
	"github.com/lessgo/lessgo"
	"github.com/lessgo/lessgoext/swagger"

	_ "github.com/lessgo/demo/BusinessAPI"
	_ "github.com/lessgo/demo/Common/Middleware"
	_ "github.com/lessgo/demo/SystemAPI"
)

func main() {
	// 开启自动api文档，false表示仅允许内网访问
	swagger.Reg(false)
	// 指定根目录URL
	lessgo.SetHome("/home")
	// 开启网络服务
	lessgo.Run()
}
