package data

import (
	"github.com/gobuffalo/packr/v2"
)

var (
	Public  *packr.Box
	Service *packr.Box
)

func init() {
	Public = packr.New("public", "../layout/public")
	Service = packr.New("service", "../layout/example")
}
