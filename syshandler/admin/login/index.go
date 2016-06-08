package login

import (
	. "github.com/lessgo/lessgo"

	"github.com/lessgo/demo/sysmodel/admin"
)

type X struct {
	Test2 string
}

var Index = ApiHandler{
	Desc:   "后台管理登录操作",
	Method: "Post",
	Params: []Param{
		{"user", "formData", true, "henry11111", "用户名"},
		{"user", "query", true, "henry22222", "用户名"},
		{"user", "path", true, "henry33333", "用户名"},
		{"password", "formData", true, "1111111111", "密码"},
		{"password", "query", true, "2222222222", "密码"},
		{"password", "path", true, "3333333333", "密码"},
	},
	HTTP200: []Result{
		{1, map[string]string{"test1": "ret1"}},
		{2, &X{Test2: "ret2"}},
		{3, []string{"ret3"}},
		{4, "ret4"},
		{5, 5},
	},
	Handler: func(c *Context) error {
		// 测试读取cookie
		id := c.CookieParam(Config.Session.SessionName)
		c.Log().Info("cookie中的%v: %#v", Config.Session.SessionName, id)

		// 测试session
		c.Log().Info("从session读取上次请求的输入: %#v", c.GetSession("info"))

		c.SetSession("info", map[string]interface{}{
			"user":     c.FormParam("user"),
			"password": c.FormParam("password"),
		})
		c.Log().Info("path用户名: %#v", c.PathParam("user"))
		c.Log().Info("query用户名: %#v", c.QueryParam("user"))
		c.Log().Info("formData用户名: %#v", c.FormParam("user"))
		c.Log().Info("path密码: %#v", c.PathParam("password"))
		c.Log().Info("query密码: %#v", c.QueryParam("password"))
		c.Log().Info("formData密码: %#v", c.FormParam("password"))

		return c.Render(200,
			"sysview/admin/login/index.tpl",
			map[string]interface{}{
				"name":       c.FormParam("user"),
				"password":   c.FormParam("password"),
				"repeatfunc": admin.Login.Repeatfunc,
			},
		)
	},
}.Reg()
