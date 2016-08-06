package admin

import (
    "time"

    "github.com/lessgo/lessgo"
)

var Index = lessgo.ApiHandler{
    Desc:   "后台首页",
    Method: "*",
    Handler: func(c *lessgo.Context) error {
        c.Log().Info("这里是后台首页,等待1s")
        c.Log().Info("获取参数A = %v", c.QueryParam("A"))
        c.Log().Info("获取参数a = %v", c.QueryParam("a"))
        time.Sleep(1e9)
        return c.JSON(200, "这里是后台首页")
    },
}.Reg()
