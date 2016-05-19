package login

import (
	. "github.com/lessgo/lessgo"

	"github.com/lessgo/demo/sysmodel/admin"
)

var IndexHandle = ApiHandler{
	Desc:   "后台管理登录操作",
	Method: "GET",
	Params: []Param{
		{"user", "path", true, "henry", "用户名"},
		{"password", "path", true, "12345678", "密码"},
	},
	Handler: func(c Context) error {
		// 测试读取cookie
		id, err := c.Request().Cookie(AppConfig.Session.SessionName)
		c.Logger().Info("cookie中的%v: %#v (%v)", AppConfig.Session.SessionName, id, err)

		// 测试session
		c.Logger().Info("从session读取上次请求的输入: %#v", c.GetSession("info"))

		c.SetSession("info", map[string]interface{}{
			"user":     c.Param("user"),
			"password": c.Param("password"),
		})

		return c.Render(200,
			"sysview/admin/login/index.tpl",
			map[string]interface{}{
				"name":       c.Param("user"),
				"password":   c.Param("password"),
				"repeatfunc": admin.Login.Repeatfunc,
			},
		)
	},
}.Reg()
