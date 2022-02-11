package conf

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bearki/gov/tool"
)

// gov version
var Version = "0.1.2"

var (
	// default GOSDK path
	GOSDKPATH string
	// go version list file path
	GOVERSIONFILE string
	// go root path
	GOROOT string
	// default GOSDK version list url
	GOSDKVERURL string
	// default GOSDK download url
	GOSDKDOWNURL string
)

// init env
func init() {
	// 从环境变量中获取SDK存放路径(环境变量优先级最高)
	goSdkPath := os.Getenv("GOSDKPATH")
	if len(goSdkPath) > 0 {
		// 赋值获取到的环境变量
		GOSDKPATH = goSdkPath
	} else {
		if runtime.GOOS == "windows" { // windows环境变量
			// Windows 环境下指定默认的GOSDKPATH为%LOCALAPPDATA%\Gov
			GOSDKPATH = filepath.Join(os.Getenv("LOCALAPPDATA"), "Gov")
		} else { // 其他系统环境变量
			// 其他环境下指定默认的GOSDKPATH
			GOSDKPATH = filepath.Join(os.Getenv("HOME"), "Gov")
		}
	}

	// 配置版本列表json文件的存放路径
	GOVERSIONFILE = filepath.Join(GOSDKPATH, "version.json")

	// 从环境变量中获取Go Root存放路径(环境变量优先级最高)
	goRootPath := os.Getenv("GOROOT")
	if len(goRootPath) > 0 {
		// 赋值获取到的环境变量
		GOROOT = goRootPath
	} else {
		if runtime.GOOS == "windows" { // windows环境变量
			// Windows 环境下指定默认的GOROOT为%LOCALAPPDATA%\Go
			GOROOT = filepath.Join(os.Getenv("LOCALAPPDATA"), "Go")
			// 将GOROOT和%GOROOT/bin%写入环境变量
			err := tool.SetRegistryValue("GOROOT", GOROOT, 1)
			if err != nil {
				tool.L.Error(err.Error())
				return
			}
			PATH, err := tool.GetRegistryValue("PATH")
			if err != nil {
				tool.L.Error(err.Error())
				return
			}
			if !strings.Contains(PATH, "%GOROOT%\\bin") {
				err = tool.SetRegistryValue("PATH", "%GOROOT%\\bin;"+PATH, 2)
				if err != nil {
					tool.L.Error(err.Error())
					return
				}
			}
		} else { // 其他系统环境变量
			// 其他环境下指定默认的GOROOT
			GOROOT = filepath.Join(os.Getenv("HOME"), "Go")
			// 写入环境变量
			envStr := fmt.Sprintf(
				"%s\n%s",
				"export GOROOT="+GOROOT,
				"export PATH=$GOROOT/bin:$PATH",
			)
			cmd := exec.Command("echo", envStr, ">>", filepath.Join(os.Getenv("HOME"), ".bashrc"))
			err := cmd.Run()
			if err != nil {
				tool.L.Error(err.Error())
				return
			}
			cmd = exec.Command("source", filepath.Join(os.Getenv("HOME"), ".bashrc"))
			err = cmd.Run()
			if err != nil {
				tool.L.Error(err.Error())
				return
			}
		}
	}

	// 从环境变量中获取SDK版本列表网址BaseUrl(环境变量优先级最高)
	goSdkVerUrl := os.Getenv("GOSDKVERURL")
	if strings.Contains(goSdkVerUrl, "http") {
		// 赋值获取到的环境变量
		GOSDKVERURL = goSdkVerUrl
	} else {
		// 赋值默认版本列表请求地址
		GOSDKVERURL = "https://qiniu.github.bearki.cn/gov/version/version.json"
	}

	// 从环境变量中获取SDK下载网址BaseUrl(环境变量优先级最高)
	goSdkDownUrl := os.Getenv("GOSDKDOWNURL")
	if strings.Contains(goSdkDownUrl, "http") {
		// 赋值获取到的环境变量
		GOSDKDOWNURL = goSdkDownUrl
	} else {
		// 赋值SDK默认下载地址
		GOSDKDOWNURL = "https://mirrors.ustc.edu.cn/golang"
	}
}
