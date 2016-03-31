package System

import (
	"github.com/lessgo/lessgo"
	"testing"
)

func TestRouter(t *testing.T) {
	for _, r := range lessgo.DynaRouterTree() {
		t.Logf("%v\n", r.Url())
		t.Logf("%#v\n\n", r)
	}
}
