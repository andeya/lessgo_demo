package router

import (
    "github.com/lessgo/lessgo"

    "github.com/lessgo/demo/bizhandler/home"
    "github.com/lessgo/demo/middleware"
)

func init() {
    lessgo.Root(
        lessgo.Leaf("/websocket", home.WebSocket, middleware.ShowHeader),
        lessgo.Branch("/home", "前台",
            lessgo.Leaf("/index", home.Index, middleware.ShowHeader),
        ).Use(middleware.Print),
    )
}
