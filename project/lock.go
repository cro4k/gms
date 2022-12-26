package project

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	lockfile  = "gms.lock.yml"
	goVersion = "1.18"
)

type LockInfo struct {
	Prefix    string   `yaml:"prefix"`
	Name      string   `yaml:"name"`
	Version   string   `yaml:"version"`
	Service   []string `yaml:"service"`
	GoVersion string   `yaml:"goversion"`
}

func loadLock() (*LockInfo, error) {
	_, err := os.Stat(lockfile)
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(lockfile)
	if err != nil {
		return nil, err
	}
	var info = new(LockInfo)
	err = yaml.Unmarshal(b, info)
	if err != nil {
		return info, err
	}
	return info, err
}

func (i *LockInfo) create(path string) error {
	var filepath = lockfile
	if path != "" {
		filepath = fmt.Sprintf("%s/%s", path, filepath)
	}
	lock, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer lock.Close()
	return yaml.NewEncoder(lock).Encode(i)
}

func (i *LockInfo) clean(path string) error {
	var filepath = lockfile
	if path != "" {
		filepath = fmt.Sprintf("%s/%s", path, filepath)
	}
	return os.Remove(filepath)
}
