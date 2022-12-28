package project

import (
	"fmt"
	"github.com/cro4k/gms/utils/fileutil"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const (
	lockfile  = ".gms/lock.yml"
	goVersion = "1.18"
)

type LockInfo struct {
	Prefix    string   `yaml:"prefix"`
	Name      string   `yaml:"name"`
	Version   string   `yaml:"version"`
	Service   []string `yaml:"service"`
	GoVersion string   `yaml:"goversion"`
	Git       bool     `yaml:"git"`
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
	gmsPath := fmt.Sprintf("%s/.gms", path)
	_ = os.MkdirAll(gmsPath, 0777)
	if err := fileutil.Hide(gmsPath); err != nil {
		log.Println(err)
	}

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

func (i *LockInfo) has(service string) bool {
	for _, v := range i.Service {
		if v == service {
			return true
		}
	}
	return false
}
