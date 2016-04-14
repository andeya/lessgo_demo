package main

import (
	"github.com/lessgo/lessgo"
	// . "github.com/lessgo/lessgo/engine/standard"
	. "github.com/lessgo/lessgo/engine/fasthttp"

	_ "github.com/lessgo/demo/BusinessAPI"
	_ "github.com/lessgo/demo/Common/Middleware"
	_ "github.com/lessgo/demo/SystemAPI"
	"github.com/lessgo/lessgoext/swagger"
)

func main() {
	swagger.Init()
	lessgo.SetHome("/home")
	lessgo.Run(WithConfig)
}
