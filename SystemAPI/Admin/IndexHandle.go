package Admin

import (
	"time"

	"github.com/lessgo/lessgo"
)

var IndexHandle = lessgo.RegHandler(lessgo.ApiHandler{
	Desc:    "后台首页",
	Methods: []string{},
	Handler: func(ctx lessgo.Context) error {
		ctx.Logger().Info("这里是后台首页,等待1s")
		ctx.Logger().Info("获取参数A = %v", ctx.QueryParam("A"))
		ctx.Logger().Info("获取参数a = %v", ctx.QueryParam("a"))
		time.Sleep(1e9)
		return ctx.JSON(200, "这里是后台首页")
	},
})
