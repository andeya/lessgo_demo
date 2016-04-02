package Business

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/Business/HomeModule"
)

func init() {
	lessgo.RootRouter(
		lessgo.SubRouter("/home", "前台", "前台应用",
			lessgo.Get("首页", "首页", HomeModule.IndexHandle).Use("/home/index中间件"),
		).Use("进入分组路由", "/home中间件"),
	)
}
