package main

import (
	"github.com/henrylee2cn/lessgo"
	"github.com/henrylee2cn/lessgoext/swagger"

	// _ "github.com/henrylee2cn/lessgoext/dbservice/gorm"
	// _ "github.com/henrylee2cn/lessgoext/dbservice/xorm"
	_ "github.com/henrylee2cn/lessgoext/dbservice/sqlx"

	_ "github.com/henrylee2cn/lessgo_demo/middleware"
	_ "github.com/henrylee2cn/lessgo_demo/router"
)

func main() {
	// 开启自动api文档，通过config/apidoc_allow.myconfig进行配置
	swagger.Reg()
	// 指定根目录URL
	lessgo.SetHome("/home")
	// 开启网络服务
	lessgo.Run()
}
