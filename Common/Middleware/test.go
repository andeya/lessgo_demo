package Middleware

import (
	"errors"

	"github.com/lessgo/lessgo"
	"github.com/lessgo/lessgo/logs"
)

var PrintWare = (&lessgo.ApiMiddleware{
	Name:          "打印测试",
	Desc:          "打印测试",
	DefaultConfig: nil,
	Middleware: func(config string) lessgo.MiddlewareFunc {
		return lessgo.WrapMiddleware(func(ctx lessgo.Context) error {
			ctx.Logger().Info("测试中间件-打印一些东西：1234567890")
			return nil
		})
	},
}).Init()

var ShowHeaderWare = (&lessgo.ApiMiddleware{
	Name:          "显示Header",
	Desc:          "显示Header测试",
	DefaultConfig: nil,
	Middleware: func(config string) lessgo.MiddlewareFunc {
		return lessgo.WrapMiddleware(func(ctx lessgo.Context) error {
			logs.Info("测试中间件-显示Header：%v", ctx.Request().Header().Keys())
			return nil
		})
	},
}).Init()

var ErrorTestWare = (&lessgo.ApiMiddleware{
	Name:          "故意报错",
	Desc:          "故意报错测试",
	DefaultConfig: nil,
	Middleware: func(config string) lessgo.MiddlewareFunc {
		return lessgo.WrapMiddleware(func(ctx lessgo.Context) error {
			ctx.Logger().Info("测试中间件-打印一些东西：1234567890")
			return errors.New("中间件故意报错 error")
		})
	},
}).Init()
