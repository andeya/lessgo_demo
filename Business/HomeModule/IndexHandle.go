package HomeModule

import (
	"github.com/lessgo/lessgo"
)

func IndexHandle(ctx lessgo.Context) error {
	ctx.String(200, "这里是首页")
	return nil
}
