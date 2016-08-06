package biz_handler

import (
	"github.com/lessgo/lessgo"
)

var GetUpload = lessgo.ApiHandler{
	Desc:   "文件上传页",
	Method: "GET",
	Handler: func(c *lessgo.Context) error {
		return c.HTML(
			200,
			`<form action="/upload" enctype="multipart/form-data" method="POST">
        <input type="file" name="test" value="文件测试">
        <input type="submit" value="上传">
    </form>`)
	},
}.Reg()

var PostUpload = lessgo.ApiHandler{
	Desc:   "文件上传",
	Method: "POST",
	Params: []lessgo.Param{
		{"test", "formData", true, nil, "文件"},
	},
	Handler: func(c *lessgo.Context) error {
		fileUrl, size, err := c.SaveFile("test", false, "/a/?")
		c.Log().Info("%v, %v, %v", fileUrl, size, err)
		return err
	},
}.Reg()
