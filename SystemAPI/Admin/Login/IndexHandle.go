package Login

import (
    . "github.com/lessgo/lessgo"
)

var IndexHandle = ApiHandler{
    Desc:    "后台管理登录操作",
    Types: []string{"GET", "OPTIONS"},
    Params: []Param{
        {
            Name:     "user",
            In:       "path",
            Required: true,
            Format:   "henry",
            Desc:     "用户名",
        },
        {
            Name:     "password",
            In:       "path",
            Required: true,
            Format:   "12345678",
            Desc:     "密码",
        },
    },
    Handler: func(ctx Context) error {
        // 测试读取cookie
        id, err := ctx.Request().Cookie(AppConfig.Session.CookieName)
        ctx.Logger().Info("cookie中的%v: %#v (%v)", AppConfig.Session.CookieName, id, err)

        // 测试session
        sess, err := ctx.Session()
        if err != nil {
            ctx.Logger().Error("%v", err)
        } else {
            ctx.Logger().Info("session中记录的上次输入: %#v", sess.Get("info"))

            err = sess.Set("info", map[string]interface{}{
                "user":     ctx.Param("user"),
                "password": ctx.Param("password"),
            })
            if err != nil {
                ctx.Logger().Error("Set session [info] failed: %v", err)
            }
        }

        return ctx.Render(200,
            "SystemView/Admin/Login/index.tpl",
            map[string]interface{}{
                "name":       ctx.Param("user"),
                "password":   ctx.Param("password"),
                "repeatfunc": repeatfunc,
            })
    },
}.Reg()
