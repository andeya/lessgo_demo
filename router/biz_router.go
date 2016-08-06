package router

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/biz_handler"
	"github.com/lessgo/demo/biz_handler/home"
	"github.com/lessgo/demo/middleware"
	extm "github.com/lessgo/lessgoext/middleware"
)

func init() {
	lessgo.Root(

		lessgo.Leaf("/upload", biz_handler.GetUpload),
		lessgo.Leaf("/upload", biz_handler.PostUpload),
		lessgo.Leaf("/md1", biz_handler.Markdown1, extm.Gzip),
		lessgo.Leaf("/md2", biz_handler.Markdown2),
		lessgo.Leaf("/websocket", biz_handler.WebSocket, middleware.ShowHeader),
		lessgo.Leaf("/checkbox", biz_handler.GetCheckbox, middleware.ShowHeader),
		lessgo.Leaf("/checkbox", biz_handler.PostCheckbox),
		lessgo.Leaf("/baidu", biz_handler.ReverseProxyBaidu),
		lessgo.Branch("/home", "前台",
			lessgo.Leaf("/index", home.Index, middleware.ShowHeader),
		).Use(middleware.Print),
	)
}
