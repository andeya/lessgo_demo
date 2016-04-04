package Home

import (
	"github.com/lessgo/lessgo"
	"github.com/lessgo/lessgo/logs"
)

func IndexHandle(ctx lessgo.Context) error {
	logs.Info("url:%v", ctx.Path())
	ctx.String(200, "这里是首页")
	return nil
}
