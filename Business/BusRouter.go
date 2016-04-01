package Business

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/Business/HomeModule"
)

func init() {
	lessgo.RootRouter(
		lessgo.SubRouter("/home", "前台", "前台应用",
			lessgo.Get("首页", "首页", HomeModule.IndexHandle),
		),
	)
}
