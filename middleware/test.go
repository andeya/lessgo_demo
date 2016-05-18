package middleware

import (
    "github.com/lessgo/lessgo"
    "github.com/lessgo/lessgo/logs"
)

var PrintWare = lessgo.ApiMiddleware{
    Name:          "打印测试",
    Desc:          "打印测试",
    DefaultConfig: nil,
    Middleware: func(config string) lessgo.MiddlewareFunc {
        return lessgo.WrapMiddleware(func(c lessgo.Context) error {
            c.Logger().Info("测试中间件-打印一些东西：1234567890")
            c.Logger().Info("param:%v(len=%v),%v(len=%v)", c.ParamNames(), len(c.ParamNames()), c.ParamValues(), len(c.ParamValues()))
            return nil
        })
    },
}.Reg()

var ShowHeaderWare = lessgo.ApiMiddleware{
    Name:          "显示Header",
    Desc:          "显示Header测试",
    DefaultConfig: nil,
    Middleware: func(c lessgo.Context) error {
        logs.Info("测试中间件-显示Header：%v", c.Request().Header)
        return nil
    },
}.Reg()
