package main

import (
	"fmt"
	"github.com/cro4k/common/args"
	"github.com/cro4k/gms/project"
	"github.com/cro4k/gms/version"
	"os"
)

//go:generate packr2
func main() {
	cmd := parse()
	if cmd.Is("h", "-help") || cmd.Cmd == "" {
		fmt.Println("gms " + version.Version)
		fmt.Println(help(cmd.Cmd))
		return
	}

	var err error
	switch cmd.Cmd {
	case "create":
		opt := project.CreateOption{
			Name:    cmd.Name,
			Prefix:  cmd.Val("p", "-prefix"),
			Service: cmd.Others,
			Go:      cmd.Val("go"),
			Git:     cmd.Is("git"),
		}
		err = project.Create(opt)
	case "init":
		err = project.Init()
	case "add":
		err = project.Add(append([]string{cmd.Name}, cmd.Others...)...)
	case "fix":
		err = project.Fix()
	case "status":
		var message string
		message, err = project.Status()
		if err == nil {
			_, _ = os.Stdout.WriteString(message)
		}
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	createHelp = "create 新建项目\n" +
		"  新建一个项目并自动初始化\n" +
		"  Usage: gms create [project] [service1] [service2]...\n" +
		"  option:\n" +
		"    -p 自定义module前缀\n" +
		"    -go 指定go版本"
	initHelp = "init 初始化当前项目\n" +
		"  根据gms.lock.yml配置初始化项目，配置文件必须存在\n" +
		"  Usage: gms init"
	addHelp = "add 添加服务\n" +
		"  向当前项目内新增服务\n" +
		"  Usage: gms add [service1] [service2]..."

	fixHelp = "fix 补全缺失服务\n" +
		"  向当前项目内，添加缺失的服务，已存在的服务将跳过创建\n" +
		"  Usage: gms fix"
	statusHelp = "status 查看状态"
)

func help(cmd string) string {

	switch cmd {
	case "create":
		return createHelp
	case "init":
		return initHelp
	case "add":
		return addHelp
	case "fix":
		return fixHelp
	case "status":
		return statusHelp
	default:
		return createHelp + "\n\n" +
			initHelp + "\n\n" +
			addHelp + "\n\n" +
			fixHelp + "\n\n" +
			statusHelp
	}
}

func parse() *Cmd {
	cmd := &Cmd{}
	kv, cmds := args.ParseArgs()
	if len(cmds) > 0 {
		cmd.Cmd = cmds[0]
	}
	if len(cmds) > 1 {
		cmd.Name = cmds[1]
	}
	if len(cmds) > 2 {
		cmd.Others = cmds[2:]
	}
	cmd.Option = kv
	return cmd
}

type Cmd struct {
	Cmd    string
	Name   string
	Others []string
	Option map[string]string
}

func (c *Cmd) Is(keys ...string) bool {
	for _, key := range keys {
		if _, ok := c.Option[key]; ok {
			return true
		}
	}
	return false
}

func (c *Cmd) Val(keys ...string) string {
	for _, key := range keys {
		if val, ok := c.Option[key]; ok {
			return val
		}
	}
	return ""
}
