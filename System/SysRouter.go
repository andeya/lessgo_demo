package System

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/System/AdminModule"
	// "github.com/lessgo/demo/System/AdminModule/CountModule"
	// "github.com/lessgo/demo/System/PowerModule"
)

func init() {
	lessgo.RootRouter(
		lessgo.SubRouter("/admin", "后台管理", "管理整个项目的运行",
			lessgo.Any("后台首页", "后台管理的首页", AdminModule.IndexHandle),
			lessgo.Get("后台登陆", "登陆管理后台", AdminModule.LoginHandle, "/:user/:password"),
			lessgo.Post("后台登陆", "登陆管理后台", AdminModule.LoginHandle, "/:user/:password"),

			// lessgo.SubRouter("/count", "访问统计", "网站访问量统计",
			// 	lessgo.Get("全局访问量", "默认显示全局访问量", CountModule.IndexHandle),
			// 	lessgo.Get("局部访问量", "依据条件显示局部访问量", CountModule.PartHandle, "/:site/:start/:end"),
			// ),
		),
	)
}
