package conf

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/bearki/gov/tool"
)

// gov version
var Version = "0.1.0"

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
			GOSDKPATH = filepath.Join("/usr", "local", "Gov")
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
			// 判断是否已经写入过
			if !strings.Contains(PATH, "%GOROOT%\\bin") {
				// 写入环境变量
				err = tool.SetRegistryValue("PATH", "%GOROOT%\\bin;"+PATH, 2)
				if err != nil {
					tool.L.Error(err.Error())
					return
				}
			}
			// 借助user32.dll库的函数刷新环境变量
			user32, err := syscall.LoadDLL("user32.dll")
			if err != nil {
				tool.L.Error(err.Error())
				return
			}
			SendMessageTimeout, err := user32.FindProc("SendMessageW")
			if err != nil {
				tool.L.Error(err.Error())
				return
			}
			var HWND_BROADCAST = 0xffff
			var WM_SETTINGCHANGE = 0x001A
			content, _ := syscall.UTF16FromString("Environment")
			_, _, _ = SendMessageTimeout.Call(
				uintptr(HWND_BROADCAST),
				uintptr(WM_SETTINGCHANGE),
				uintptr(0),
				uintptr(unsafe.Pointer(&content[0])),
			)
		} else { // 其他系统环境变量
			// 其他环境下指定默认的GOROOT
			GOROOT = filepath.Join("/usr", "local", "Go")
			// 写入环境变量
			envStr := fmt.Sprintf(
				"%s\n%s",
				"export GOROOT="+GOROOT,
				"export PATH=$GOROOT/bin:$PATH",
			)
			cmd := exec.Command("echo", "/etc/profile", "<<", envStr)
			err := cmd.Run()
			if err != nil {
				tool.L.Error(err.Error())
				return
			}
			cmd = exec.Command("source", "/etc/profile")
			err = cmd.Run()
			if err != nil {
				tool.L.Error(err.Error())
				return
			}
		}
	}

	// 从环境变量中获取SDK版本列表网址BaseUrl(环境变量优先级最高)
	goSdkVerUrl := os.Getenv("GOSDKVERURL")
	if len(goSdkVerUrl) > 0 {
		// 赋值获取到的环境变量
		GOSDKVERURL = goSdkVerUrl
	} else {
		// 赋值默认版本列表请求地址
		GOSDKVERURL = "https://qiniu.github.bearki.cn/gov/version/version.json"
	}

	// 从环境变量中获取SDK下载网址BaseUrl(环境变量优先级最高)
	goSdkDownUrl := os.Getenv("GOSDKDOWNURL")
	if len(goSdkDownUrl) > 0 {
		// 赋值获取到的环境变量
		GOSDKDOWNURL = goSdkDownUrl
	} else {
		// 赋值SDK默认下载地址
		GOSDKDOWNURL = "https://mirrors.ustc.edu.cn/golang"
	}
}
