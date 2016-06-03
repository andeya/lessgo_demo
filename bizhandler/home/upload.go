package home

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
	Handler: func(c *lessgo.Context) error {
		fname, size, err := c.SaveFile("test", false, "/a/?")
		c.Log().Info("%v, %v, %v", fname, size, err)
		return err
	},
}.Reg()
