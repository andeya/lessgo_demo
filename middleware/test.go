package middleware

import (
	"github.com/henrylee2cn/lessgo"
)

var Print = lessgo.ApiMiddleware{
	Name:   "打印测试",
	Desc:   "打印测试1",
	Config: nil,
	Middleware: func(confObject interface{}) lessgo.MiddlewareFunc {
		return lessgo.WrapMiddleware(func(c *lessgo.Context) error {
			c.Log().Info("测试中间件-打印一些东西：1234567890")
			c.Log().Info("param:%v(len=%v),%v(len=%v)", c.PathParamKeys(), len(c.PathParamKeys()), c.PathParamValues(), len(c.PathParamValues()))
			return nil
		})
	},
}.Reg()

var ShowHeader = lessgo.ApiMiddleware{
	Name:   "显示Header",
	Desc:   "显示Header测试2",
	Config: nil,
	Middleware: func(c *lessgo.Context) error {
		c.Log().Info("测试中间件-显示Header：%v", c.Request().Header)
		return nil
	},
}.Reg()

var Param1 = lessgo.ApiMiddleware{
	Name: "测试参数",
	Desc: "测试参数",
	Params: []lessgo.Param{
		{"pathParam", "path", true, "", "来自中间件"},
		{"queryParam", "query", true, "", "来自中间件"},
		{"formParam", "formData", false, "", "来自中间件"},
	},
	Config: nil,
	Middleware: func(c *lessgo.Context) error {
		c.Log().Info("pathParam: %v", c.PathParam("pathParam"))
		c.Log().Info("queryParam: %v", c.QueryParam("queryParam"))
		c.Log().Info("formParam: %v", c.FormParam("formParam"))
		return nil
	},
}.Reg()

var Param2 = lessgo.ApiMiddleware{
	Name: "测试参数",
	Desc: "测试参数",
	Params: []lessgo.Param{
		{"user", "formData", false, "henry11111", "用户名"},
		{"user2", "formData", false, "henry11111", "用户名"},
		{"password", "formData", true, "1111111111", "密码"},
		{"password2", "formData", true, "1111111111", "密码"},
	},
	Config: nil,
	Middleware: func(c *lessgo.Context) error {
		c.Log().Info("user: %v", c.FormParam("user"))
		c.Log().Info("user2: %v", c.FormParam("user"))
		c.Log().Info("password: %v", c.FormParam("password"))
		c.Log().Info("password2: %v", c.FormParam("password"))
		return nil
	},
}.Reg()
