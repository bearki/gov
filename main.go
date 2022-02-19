package main

import (
	"fmt"

	"github.com/bearki/gov/cmd"
	"github.com/bearki/gov/conf"
	"github.com/bearki/gov/tool"
	"github.com/gookit/color"
)

const logo = `
 ________  ________  ___      ___ 
|\   ____\|\   __  \|\  \    /  /|
\ \  \___|\ \  \|\  \ \  \  /  / /
 \ \  \  __\ \  \\\  \ \  \/  / / 
  \ \  \|\  \ \  \\\  \ \    / /  
   \ \_______\ \_______\ \__/ /   
    \|_______|\|_______|\|__|/    

`
const welcome = `Welcome to Gov, an awesome Golang language version switcher.`

func main() {
	fmt.Println(color.Magenta.Sprint(logo))
	tool.L.Warn("Gov Version %s", conf.Version)
	tool.L.Trace(welcome)
	tool.L.Info(tool.StartLine)
	defer tool.L.Info(tool.EndLine)
	err := conf.Init()
	if err != nil {
		return
	}
	cmd.Execute()
}
