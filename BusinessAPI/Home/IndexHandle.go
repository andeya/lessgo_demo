package Home

import (
	"github.com/lessgo/lessgo"
)

func IndexHandle(ctx lessgo.Context) error {
	return ctx.Render(
		200,
		"BusinessView/Home/index.tpl",
		map[string]interface{}{
			"title":   lessgo.NAME,
			"version": lessgo.VERSION,
			"content": "Welcome To Lessgo",
		})
}
