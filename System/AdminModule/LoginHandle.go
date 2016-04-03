package AdminModule

import (
	"github.com/lessgo/lessgo"
)

func LoginHandle(ctx lessgo.Context) error {
	return ctx.Render(200,
		"System/AdminModule/View/login.html",
		map[string]interface{}{
			"name":     ctx.Param("user"),
			"password": ctx.Param("password"),
		})
}
