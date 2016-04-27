package Login

import (
	. "github.com/lessgo/lessgo"
)

var IndexHandle = DescHandler{
	Desc:     "后台管理登录操作",
	Produces: []string{"application/html"},
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
		sess, err := ctx.SessionStart()
		if err != nil {
			ctx.Logger().Error("%v", err)
		} else {
			err = sess.Set("user", ctx.Param("user"))
			if err != nil {
				ctx.Logger().Error("Set session [user] failed: %v", err)
			}
			err = sess.Set("password", ctx.Param("password"))
			if err != nil {
				ctx.Logger().Error("Set session [user] failed: %v", err)
			}
			ctx.Logger().Info("session user: %v", sess.Get("user"))
			ctx.Logger().Info("session password: %v", sess.Get("password"))
		}

		return ctx.Render(200,
			"SystemView/Admin/Login/index.tpl",
			map[string]interface{}{
				"name":       ctx.Param("user"),
				"password":   ctx.Param("password"),
				"repeatfunc": repeatfunc,
			})
	},
}
