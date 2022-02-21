package main

import (
	"fmt"
	"os"

	"github.com/bearki/gov/cmd"
	"github.com/bearki/gov/conf"
	"github.com/bearki/gov/tool"
	"github.com/gookit/color"
)

const logo = `
 $$$$$$\   $$$$$$\  $$\    $$\
$$  __$$\ $$  __$$\ $$ |   $$ |
$$ /  \__|$$ /  $$ |$$ |   $$ |
$$ |$$$$\ $$ |  $$ |\$$\  $$  |
$$ |\_$$ |$$ |  $$ | \$$\$$  /
$$ |  $$ |$$ |  $$ |  \$$$  /
\$$$$$$  | $$$$$$  |   \$  /
 \______/  \______/     \_/

 `

const welcome = `Welcome to Gov, an awesome Golang language version switcher.`

func main() {
	// 仅在无指令或指令为help或指令不存在时打印LOGO和版本
	cmdMap := cmd.GetCmdNameMap()
	if len(os.Args) < 2 {
		printLogoVersion()
	} else if _, ok := cmdMap[os.Args[1]]; !ok {
		printLogoVersion()
	}

	// 打印任务线
	tool.L.Info(tool.StartLine)
	defer tool.L.Info(tool.EndLine)

	// 初始化一下配置
	err := conf.Init()
	if err != nil {
		return
	}

	// 执行命令
	cmd.Execute()
}

// printLogoVersion 打印LOGO和版本
func printLogoVersion() {
	fmt.Println(color.Magenta.Sprint(logo))
	tool.L.Warn("Gov Version %s", conf.Version)
	tool.L.Trace(welcome)
}
