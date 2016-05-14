package Login

import (
	. "github.com/lessgo/lessgo"
)

var IndexHandle = ApiHandler{
	Desc:   "后台管理登录操作",
	Method: "GET",
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
		id, err := ctx.Request().Cookie(AppConfig.Session.SessionName)
		ctx.Logger().Info("cookie中的%v: %#v (%v)", AppConfig.Session.SessionName, id, err)

		// 测试session
		ctx.Logger().Info("从session读取上次请求的输入: %#v", ctx.GetSession("info"))

		ctx.SetSession("info", map[string]interface{}{
			"user":     ctx.Param("user"),
			"password": ctx.Param("password"),
		})

		return ctx.Render(200,
			"SystemView/Admin/Login/index.tpl",
			map[string]interface{}{
				"name":       ctx.Param("user"),
				"password":   ctx.Param("password"),
				"repeatfunc": repeatfunc,
			},
		)
	},
}.Reg()
