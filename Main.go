package main

import (
	"github.com/lessgo/lessgo"
	// "github.com/lessgo/lessgo/engine/standard"
	"github.com/lessgo/lessgo/engine/fasthttp"

	_ "github.com/lessgo/demo/Business"
	_ "github.com/lessgo/demo/Common/Middleware"
	_ "github.com/lessgo/demo/System"
)

func main() {
	// lessgo.Run(standard.NewFromConfig)
	lessgo.Run(fasthttp.NewFromConfig)
}
