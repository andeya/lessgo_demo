package BusinessAPI

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/BusinessAPI/Home"
)

func init() {
	lessgo.RootRouter(
		lessgo.SubRouter("/home", "前台", "前台应用",
			lessgo.Get("首页", "首页", Home.IndexHandle).Use("显示Header"),
		).Use("打印一些东西"),
	)
}
