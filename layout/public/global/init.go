package global

import (
	"flag"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	c          *Config
	configPath string
)

func init() {
	flag.StringVar(&configPath, "c", "conf.yml", "config path\nremote http(s) path or local file path")
	flag.Parsed()
	c = new(Config)
	if err := loadConfig(c, configPath); err != nil {
		logrus.Fatal(err)
	}
}

func loadConfig(v interface{}, path string) error {
	var body io.Reader
	if strings.HasPrefix(path, "http") {
		if resp, err := http.Get(path); err != nil {
			return err
		} else {
			defer resp.Body.Close()
			body = resp.Body
		}
	} else {
		if fi, err := os.Open(path); err != nil {
			return err
		} else {
			defer fi.Close()
			body = fi
		}
	}
	err := yaml.NewDecoder(body).Decode(v)
	return err
}

func LoadConfig(v interface{}) error {
	return loadConfig(v, configPath)
}

func C() *Config {
	return c
}
