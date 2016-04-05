package Login

import (
	"github.com/lessgo/lessgo"
)

func IndexHandle(ctx lessgo.Context) error {
	return ctx.Render(200,
		"SystemView/Admin/Login/index.tpl",
		map[string]interface{}{
			"name":       ctx.Param("user"),
			"password":   ctx.Param("password"),
			"repeatfunc": repeatfunc,
		})
}
