package SystemAPI

import (
	. "github.com/lessgo/lessgo"

	"github.com/lessgo/demo/SystemAPI/Admin"
	"github.com/lessgo/demo/SystemAPI/Admin/Login"
)

func init() {
	RootRouter(
		SubRouter("/admin", "后台管理",
			Any("index", "后台首页", Admin.IndexHandle),
			SubRouter("/login", "后台登陆",
				Get(":user/:password", "后台登陆", Login.IndexHandle),
				Post(":user/:password", "后台登陆", Login.IndexHandle),
			),
		),
	)
}
