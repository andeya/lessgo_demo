package Middleware

import (
	"errors"

	"github.com/lessgo/lessgo"
	"github.com/lessgo/lessgo/logs"
)

func init() {
	lessgo.RegMiddleware("进入分组路由", "描述可以省略", func(ctx lessgo.Context) error {
		logs.Info("进入分组路由")
		return nil
	})
	lessgo.RegMiddleware("/home中间件", "描述可以省略", func(ctx lessgo.Context) error {
		logs.Info("/home中间件")
		return nil
	})
	lessgo.RegMiddleware("/home/index中间件", "描述可以省略", func(ctx lessgo.Context) error {
		logs.Info("/home/index中间件")
		return errors.New("/home/index中间件 error")
	})
}
