package biz_handler

import (
	"github.com/lessgo/lessgo"
)

var GetCheckbox = lessgo.ApiHandler{
	Desc:   "多选测试",
	Method: "GET",
	Params: []lessgo.Param{
		{"Fruit", "query", true, "a", "多选"},
	},
	Handler: func(c *lessgo.Context) error {
		c.Log().Warn(c.QueryParam("Fruit"))
		c.Log().Warn("%#v", c.QueryParams("Fruit"))
		c.Log().Warn("%#v", c.QueryValues())
		return c.HTML(
			200,
			`<form action="/checkbox" method="post"> 
您喜欢的水果？<br /><br /> 
<label><input name="Fruit" type="checkbox" value="a" />苹果 </label> 
<label><input name="Fruit" type="checkbox" value="b" />桃子 </label> 
<label><input name="Fruit" type="checkbox" value="c" />香蕉 </label> 
<label><input name="Fruit" type="checkbox" value="d" />梨 </label> 
<input type="submit">
</form> `)
	},
}.Reg()

var PostCheckbox = lessgo.ApiHandler{
	Desc:   "多选测试",
	Method: "POST",
	Params: []lessgo.Param{
		{"Fruit", "formData", true, "a", "多选"},
	},
	Handler: func(c *lessgo.Context) error {
		c.Log().Warn(c.FormParam("Fruit"))
		c.Log().Warn("%#v", c.FormParams("Fruit"))
		c.Log().Warn("%#v", c.FormValues())
		return nil
	},
}.Reg()
