package Middleware

import (
	"errors"

	"github.com/lessgo/lessgo"
	"github.com/lessgo/lessgo/logs"
	"github.com/lessgo/lessgoext/middleware"
)

func init() {
	lessgo.RegMiddleware("打印一些东西", "描述可以省略", func(ctx lessgo.Context) error {
		ctx.Logger().Info("打印一些东西：1234567890")
		return nil
	})
	lessgo.RegMiddleware("显示Header", "描述可以省略", func(ctx lessgo.Context) error {
		logs.Info("%v", ctx.Request().Header().Keys())
		return nil
	})
	lessgo.RegMiddleware("故意报错", "描述可以省略", func(ctx lessgo.Context) error {
		logs.Info("中间件故意报错")
		return errors.New("中间件故意报错 error")
	})
	lessgo.RegMiddleware("static", "serves static content from ./Uploads.", middleware.Static("Uploads"))
}
