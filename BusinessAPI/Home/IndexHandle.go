package Home

import (
	. "github.com/lessgo/lessgo"
)

var IndexHandle = &DescHandler{

	Handler: func(ctx Context) error {
		return ctx.Render(
			200,
			"BusinessView/Home/index.tpl",
			map[string]interface{}{
				"title":   NAME,
				"version": VERSION,
				"content": "Welcome To Lessgo",
			})
	},
}
