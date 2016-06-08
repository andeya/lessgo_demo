package router

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/middleware"
	"github.com/lessgo/demo/syshandler/admin"
	"github.com/lessgo/demo/syshandler/admin/login"
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
