package main

import (
	"github.com/lessgo/lessgo"
	"github.com/lessgo/lessgo/engine/standard"
)

func main() {
	lessgo.Run(standard.NewFromConfig)
}
