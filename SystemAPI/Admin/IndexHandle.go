package Admin

import (
	"time"

	"github.com/lessgo/lessgo"
)

func IndexHandle(ctx lessgo.Context) error {
	ctx.Logger().Info("这里是后台首页,等待1s")
	ctx.Logger().Info("获取参数A = %v", ctx.QueryParam("A"))
	ctx.Logger().Info("获取参数a = %v", ctx.QueryParam("a"))
	time.Sleep(1e9)
	return ctx.String(200, "这里是后台首页")
}
