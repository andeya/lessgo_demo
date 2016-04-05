package SystemAPI

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/SystemAPI/Admin"
	"github.com/lessgo/demo/SystemAPI/Admin/Login"
)

func init() {
	lessgo.RootRouter(
		lessgo.SubRouter("/admin", "后台管理", "管理整个项目的运行",
			lessgo.Any("后台首页", "后台管理的首页", Admin.IndexHandle),
			lessgo.Get("后台登陆", "登陆管理后台", Login.LoginHandle, "/:user/:password"),
			lessgo.Post("后台登陆", "登陆管理后台", Login.LoginHandle, "/:user/:password"),
		),
	)
}
