package SystemAPI

import (
	"github.com/lessgo/lessgo"
	"testing"
)

func TestRouter(t *testing.T) {
	for _, r := range lessgo.DynaRouterTree() {
		// t.Logf("%v\n", r.Url)
		t.Logf("%#v\n\n", r)
	}
	lessgo.ResetRealRoute()
	t.Logf("%#v\n\n", lessgo.DefLessgo.Echo.Routes())
	lessgo.DelRouter("c920a17fdc112cd0fd9a741ff6e3c7ae")
	lessgo.ResetRealRoute()
	t.Logf("%#v\n\n", lessgo.DefLessgo.Echo.Routes())
}
