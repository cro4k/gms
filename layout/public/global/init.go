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
	//logrus.SetReportCaller(true)

	flag.StringVar(&configPath, "c", "conf.yml", "config path\nremote http(s) path or local file path")
	flag.Parsed()
	var err error
	c, err = loadConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}
}

func loadConfig(path string) (*Config, error) {
	var body io.Reader
	if strings.HasPrefix(path, "http") {
		if resp, err := http.Get(path); err != nil {
			return nil, err
		} else {
			defer resp.Body.Close()
			body = resp.Body
		}
	} else {
		if fi, err := os.Open(path); err != nil {
			return nil, err
		} else {
			defer fi.Close()
			body = fi
		}
	}
	var config = new(Config)
	err := yaml.NewDecoder(body).Decode(config)
	return config, err
}

func C() *Config {
	return c
}
