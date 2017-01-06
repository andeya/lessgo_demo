package router

import (
	"github.com/henrylee2cn/lessgo"

	"github.com/henrylee2cn/lessgo_demo/middleware"
	"github.com/henrylee2cn/lessgo_demo/sys_handler/admin"
	"github.com/henrylee2cn/lessgo_demo/sys_handler/admin/login"
)

func init() {
	lessgo.Root(
		lessgo.Branch("/admin", "后台管理",
			lessgo.Leaf("/index", admin.Index),
			lessgo.Branch("/login", "后台登陆",
				lessgo.Leaf("/", login.Index),
			).Use(middleware.Param2),
		).Use(middleware.Param1),
	)
}
