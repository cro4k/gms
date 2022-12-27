package config

import (
	"github.com/cro4k/gms/layout/public/global"
	"github.com/sirupsen/logrus"
)

var (
	c *Config
)

func init() {
	c = new(Config)
	if err := global.LoadConfig(c); err != nil {
		logrus.Fatal(err)
	}
}

func C() *Config {
	return c
}
