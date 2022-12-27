package project

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cro4k/gms/data"
	"github.com/cro4k/gms/version"
	"github.com/gobuffalo/packr/v2"
	"os"
	"path/filepath"
	"strings"
)

const (
	layout        = "github.com/cro4k/gms/layout"
	publicLayout  = layout + "/public"
	serviceLayout = layout + "/example"

	PatternProjectName = "{{project}}"
	PatternServiceName = "{{service}}"
)

type CreateOption struct {
	Name    string
	Prefix  string
	Service []string
	Go      string
}

func (c CreateOption) lockfile() *LockInfo {
	info := &LockInfo{
		Prefix:    c.Prefix,
		Name:      c.Name,
		Version:   version.Version,
		Service:   c.Service,
		GoVersion: c.Go,
	}
	if info.Name == "" {
		info.Name = "example"
	}
	if info.GoVersion == "" {
		info.GoVersion = goVersion
	}
	if info.Prefix == "" {
		info.Prefix = c.Name
	}
	return info
}

func Create(options ...CreateOption) error {
	var option = CreateOption{Name: "example"}
	if len(options) > 0 {
		option = options[0]
	}
	if _, err := os.Stat(option.Name); os.IsExist(err) {
		return errors.New("project '" + option.Name + "' has been existed.")
	}
	if err := os.MkdirAll(option.Name, 0744); err != nil {
		return err
	}

	lock := option.lockfile()
	if err := lock.create(option.Name); err != nil {
		return err
	} else {
		return initProject(option.Name, lock)
	}
}

func Init() error {
	lock, err := loadLock()
	if err != nil {
		return err
	}
	return initProject("", lock)
}

func initProject(path string, lock *LockInfo) error {
	if path == "" {
		path = "."
	}
	if err := os.MkdirAll(fmt.Sprintf("%s/public", path), 0777); err != nil {
		return err
	}
	if err := createPublic(path, lock); err != nil {
		return err
	}
	for _, service := range lock.Service {
		if err := createService(path, service, lock); err != nil {
			return err
		}
	}
	return nil
}

func createPublic(path string, lock *LockInfo) error {
	return create(data.Public, path, "public", publicLayout, false, lock)
}

func createService(path string, service string, lock *LockInfo) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", path, service)); os.IsExist(err) {
		return errors.New("service '" + service + "' has been existed.")
	}
	return create(data.Service, path, service, serviceLayout, true, lock)
}

func create(resources *packr.Box, path string, name string, layout string, replacePublic bool, lock *LockInfo) error {
	os.MkdirAll(fmt.Sprintf("%s/%s", path, name), 0777)
	module := strings.Trim(fmt.Sprintf("%s/%s", lock.Prefix, name), "/")
	publicModule := strings.Trim(fmt.Sprintf("%s/public", lock.Prefix), "/")
	for _, filename := range resources.List() {
		content, err := resources.Find(filename)
		if err != nil {
			return err
		}
		content = bytes.ReplaceAll(content, []byte(layout), []byte(module))
		content = bytes.ReplaceAll(content, []byte(publicLayout), []byte(publicModule))
		content = bytes.ReplaceAll(content, []byte(PatternProjectName), []byte(lock.Name))
		content = bytes.ReplaceAll(content, []byte(PatternServiceName), []byte(name))
		output := strings.ReplaceAll(fmt.Sprintf("%s/%s/%s", path, name, filename), "\\", "/")
		os.MkdirAll(filepath.Dir(output), 0777)
		if err := os.WriteFile(output, content, 0644); err != nil {
			return err
		}
	}
	mod := fmt.Sprintf("%s/%s/go.mod", path, name)

	modContent := fmt.Sprintf(gomod, module, lock.GoVersion, lock.Prefix)
	if replacePublic {
		modContent += "\n" + fmt.Sprintf("replace %s v0.0.0 => ../public\n", publicModule)
	}
	err := os.WriteFile(mod, []byte(modContent), 0644)
	return err
}

var gomod = `module %s

go %s

require (
	%s/public v0.0.0
	github.com/cro4k/common v0.0.6
	github.com/cro4k/doc v0.0.9
	github.com/cro4k/ginx v0.0.3
	github.com/cro4k/micro v0.0.2
	github.com/gin-gonic/gin v1.8.1
	github.com/go-gormigrate/gormigrate/v2 v2.0.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/google/uuid v1.3.0
	github.com/sirupsen/logrus v1.9.0
	go.etcd.io/etcd/client/v3 v3.5.6
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)
`

func Add(name ...string) error {
	lock, err := loadLock()
	if err != nil {
		return err
	}
	for _, v := range name {
		if err := add(lock, v); err != nil {
			return err
		}
	}
	return nil
}

func add(lock *LockInfo, name string) error {
	if err := createService(".", name, lock); err != nil {
		return err
	}
	lock.Service = append(lock.Service, name)
	_ = lock.clean(".")
	return lock.create(".")
}

func Fix() error {
	lock, err := loadLock()
	if err != nil {
		return err
	}

	if _, err := os.Stat("public"); os.IsNotExist(err) {
		if err := createPublic(".", lock); err != nil {
			return err
		}
	}

	for _, service := range lock.Service {
		if _, err := os.Stat(service); os.IsExist(err) {
			continue
		}
		if err := createService(".", service, lock); err != nil {
			return err
		}
	}
	return nil
}
