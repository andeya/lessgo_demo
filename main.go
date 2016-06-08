package main

import (
	"github.com/lessgo/lessgo"
	"github.com/lessgo/lessgoext/swagger"

	// _ "github.com/lessgo/lessgoext/dbservice/gorm"
	// _ "github.com/lessgo/lessgoext/dbservice/xorm"

	_ "github.com/lessgo/demo/middleware"
	_ "github.com/lessgo/demo/router"
)

func main() {
	// 开启自动api文档，false表示仅允许内网访问
	swagger.Reg(false)
	// 指定根目录URL
	lessgo.SetHome("/home")
	// 开启网络服务
	lessgo.Run()
}
