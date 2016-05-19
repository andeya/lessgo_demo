package admin

import (
    "strings"
)

type login struct{}

var Login = login{}

func (_ login) Repeatfunc(s string, count int) string {
    return strings.Repeat(s, count)
}
