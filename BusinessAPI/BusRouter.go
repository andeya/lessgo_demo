package BusinessAPI

import (
    "github.com/lessgo/lessgo"

    "github.com/lessgo/demo/BusinessAPI/Home"
    "github.com/lessgo/demo/Common/Middleware"
)

func init() {
    lessgo.Root(
        lessgo.Leaf("/websocket", Home.WebSocketHandle, Middleware.ShowHeaderWare),
        lessgo.Branch("/home", "前台",
            lessgo.Leaf("/index", Home.IndexHandle, Middleware.ShowHeaderWare),
        ).Use(Middleware.PrintWare),
    )
}
