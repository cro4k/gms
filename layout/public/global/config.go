package global

import "fmt"

const (
	Debug   = "debug"
	Develop = "develop"
	Produce = "produce"
)

type DBConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8&loc=Asia%%2FShanghai",
		c.User, c.Pass, c.Host, c.Port, c.Name,
	)
}

func (c *DBConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type EtcdConfig struct {
	Endpoints []string `yaml:"endpoints"`
}

type Config struct {
	Env   string     `yaml:"env"`
	DB    DBConfig   `yaml:"db"`
	Redis DBConfig   `yaml:"redis"`
	Etcd  EtcdConfig `yaml:"etcd"`
}

func (c *Config) Develop() bool { return c.Env == Develop }
func (c *Config) Debug() bool   { return c.Env == Debug }
func (c *Config) Produce() bool { return c.Env == Produce }

func (c *Config) KEY() []byte {
	switch c.Env {
	case Produce:
		return []byte{234, 211, 45, 47, 242, 70, 173, 220, 26, 163, 56, 176, 128, 191, 166, 136}
	default:
		return []byte{41, 66, 164, 19, 56, 145, 63, 42, 118, 59, 163, 61, 180, 27, 45, 43}
	}
}
