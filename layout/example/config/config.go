package config

import "github.com/cro4k/gms/layout/public/global"

const (
	APIPort = 7788
	RPCPort = 7789
	RPCHost = "{{service}}"
)

type Config struct {
	global.Config
}