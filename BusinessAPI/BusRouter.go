package BusinessAPI

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/BusinessAPI/Home"
)

func init() {
	lessgo.RootRouter(
		lessgo.SubRouter("/", "前台",
			lessgo.Get("index", "首页", Home.IndexHandle, "显示Header"),
		).Use("打印一些东西"),
	)
}
