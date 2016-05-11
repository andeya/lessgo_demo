package main

import (
    "github.com/lessgo/lessgo"
    "github.com/lessgo/lessgoext/swagger"

    _ "github.com/lessgo/demo/BusinessAPI"
    _ "github.com/lessgo/demo/Common/Middleware"
    _ "github.com/lessgo/demo/SystemAPI"
)

func main() {
    swagger.Init()
    lessgo.SetHome("/home")
    lessgo.Run()
}
