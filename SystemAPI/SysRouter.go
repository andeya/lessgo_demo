package SystemAPI

import (
    "github.com/lessgo/lessgo"

    "github.com/lessgo/demo/SystemAPI/Admin"
    "github.com/lessgo/demo/SystemAPI/Admin/Login"
)

func init() {
    lessgo.Root(
        lessgo.Branch("/admin", "后台管理",
            lessgo.Leaf("/index", Admin.IndexHandle),
            lessgo.Branch("/login", "后台登陆",
                lessgo.Leaf("/", Login.IndexHandle),
            ),
        ),
    )
}
