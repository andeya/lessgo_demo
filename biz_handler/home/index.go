package home

import (
	"github.com/henrylee2cn/lessgo"
)

var Index = lessgo.ApiHandler{
	Desc:   "首页",
	Method: "GET",
	Handler: func(c *lessgo.Context) error {
		return c.Render(
			200,
			"biz_view/home/index.tpl",
			map[string]interface{}{
				"title":   lessgo.NAME,
				"version": lessgo.VERSION,
				"content": "Welcome To Lessgo",
			})
	},
}.Reg()
