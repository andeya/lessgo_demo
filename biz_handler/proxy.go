package biz_handler

import (
	"github.com/henrylee2cn/lessgo"
)

var ReverseProxyBaidu = lessgo.ApiHandler{
	Desc:   "反向代理百度",
	Method: "GET",
	Params: []lessgo.Param{
		{"wd", "query", true, "lessgo", "搜索关键词"},
	},
	Handler: func(c *lessgo.Context) error {
		return c.ReverseProxy("https://www.baidu.com/s?ie=UTF-8", false)
	},
}.Reg()
