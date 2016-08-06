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
	// 开启自动api文档
	// 参数为true表示自定义允许访问的ip前缀
	// 参数为false表示只允许局域网访问
	swagger.Reg(true)
	// 指定根目录URL
	lessgo.SetHome("/home")
	// 开启网络服务
	lessgo.Run()
}
