package Home

import (
	"github.com/lessgo/lessgo"
)

var IndexHandle = lessgo.ApiHandler{
	Desc:   "首页",
	Method: "GET",
	Handler: func(ctx lessgo.Context) error {
		return ctx.Render(
			200,
			"BusinessView/Home/index.tpl",
			map[string]interface{}{
				"title":   lessgo.NAME,
				"version": lessgo.VERSION,
				"content": "Welcome To Lessgo",
			})
	},
}.Reg()
