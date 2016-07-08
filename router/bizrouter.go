package router

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/demo/bizhandler"
	"github.com/lessgo/demo/bizhandler/home"
	"github.com/lessgo/demo/middleware"
	extm "github.com/lessgo/lessgoext/middleware"
)

func init() {
	lessgo.Root(

		lessgo.Leaf("/upload", bizhandler.GetUpload),
		lessgo.Leaf("/upload", bizhandler.PostUpload),
		lessgo.Leaf("/md1", bizhandler.Markdown1, extm.Gzip),
		lessgo.Leaf("/md2", bizhandler.Markdown2),
		lessgo.Leaf("/websocket", bizhandler.WebSocket, middleware.ShowHeader),
		lessgo.Leaf("/checkbox", bizhandler.GetCheckbox, middleware.ShowHeader),
		lessgo.Leaf("/checkbox", bizhandler.PostCheckbox),
		lessgo.Leaf("/baidu", bizhandler.ReverseProxyBaidu),
		lessgo.Branch("/home", "前台",
			lessgo.Leaf("/index", home.Index, middleware.ShowHeader),
		).Use(middleware.Print),
	)
}
