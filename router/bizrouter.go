package router

import (
    "github.com/lessgo/lessgo"

    "github.com/lessgo/demo/bizhandler/home"
    "github.com/lessgo/demo/middleware"
)

func init() {
    lessgo.Root(
        lessgo.Leaf("/websocket", home.WebSocketHandle, middleware.ShowHeaderWare),
        lessgo.Branch("/home", "前台",
            lessgo.Leaf("/index", home.IndexHandle, middleware.ShowHeaderWare),
        ).Use(middleware.PrintWare),
    )
}
