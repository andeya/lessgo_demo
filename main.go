package main

import (
	"github.com/lessgo/lessgo"
	"github.com/lessgo/lessgoext/swagger"

	// _ "github.com/lessgo/lessgoext/dbservice/gorm"
	// _ "github.com/lessgo/lessgoext/dbservice/xorm"
	_ "github.com/lessgo/lessgoext/dbservice/sqlx"

	_ "github.com/lessgo/demo/middleware"
	_ "github.com/lessgo/demo/router"
)

func main() {
	// 开启自动api文档，通过config/apidoc_allow.myconfig进行配置
	swagger.Reg()
	// 指定根目录URL
	lessgo.SetHome("/home")
	// 开启网络服务
	lessgo.Run()
}
