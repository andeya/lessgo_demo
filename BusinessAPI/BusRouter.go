package BusinessAPI

import (
	"github.com/lessgo/lessgo"

	"github.com/lessgo/lessgoext/middleware"

	"github.com/lessgo/demo/BusinessAPI/Home"
	"github.com/lessgo/demo/Common/Middleware"
)

func init() {
	lessgo.Root(
		lessgo.Branch("/home", "前台",
			lessgo.Leaf("/index", Home.IndexHandle, middleware.OnlyLANAccessWare, Middleware.ShowHeaderWare),
		).Use(Middleware.PrintWare),
	)
}
