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
			Format:   "",
			Desc:     "用户名",
		},
		{
			Name:     "password",
			In:       "path",
			Required: true,
			Format:   "",
			Desc:     "密码",
		},
	},
	Handler: func(ctx Context) error {
		return ctx.Render(200,
			"SystemView/Admin/Login/index.tpl",
			map[string]interface{}{
				"name":       ctx.Param("user"),
				"password":   ctx.Param("password"),
				"repeatfunc": repeatfunc,
			})
	},
}
