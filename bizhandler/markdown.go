package bizhandler

import (
	"github.com/lessgo/lessgo"
)

var Markdown1 = lessgo.ApiHandler{
	Desc:   "Markdown测试(自动生成目录)",
	Method: "GET",
	Handler: func(c *lessgo.Context) error {
		return c.Markdown("bizview/test.md", true)
	},
}.Reg()

var Markdown2 = lessgo.ApiHandler{
	Desc:   "Markdown测试(不含目录)",
	Method: "GET",
	Handler: func(c *lessgo.Context) error {
		return c.Markdown("bizview/test.md")
	},
}.Reg()
