package BusinessAPI

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/BusinessAPI/Home"
)

func init() {
	lessgo.RootRouter(
		lessgo.SubRouter("/", "前台", "前台应用",
			lessgo.Get("首页", "首页", Home.IndexHandle, "?a").Use("/home/index中间件"),
		).Use("进入分组路由", "/home中间件"),
	)
}
