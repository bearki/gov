package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bearki/gov/tool"
)

// gov version
var Version = "0.1.7"

// env filed define
const (
	ENVGOSDKPATH     = "GOSDKPATH"
	ENVGOVERSIONFILE = "GOVERSIONFILE"
	ENVGOROOT        = "GOROOT"
	ENVGOSDKVERURL   = "GOSDKVERURL"
	ENVGOSDKDOWNURL  = "GOSDKDOWNURL"
)

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
func Init() error {
	// 从环境变量中获取SDK存放路径(环境变量优先级最高)
	goSdkPath := os.Getenv(ENVGOSDKPATH)
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
	goRootPath := os.Getenv(ENVGOROOT)
	if len(goRootPath) > 0 {
		// 赋值获取到的环境变量
		GOROOT = goRootPath
	} else {
		if runtime.GOOS == "windows" { // windows环境变量
			// Windows 环境下指定默认的GOROOT为%LOCALAPPDATA%\Go
			GOROOT = filepath.Join(os.Getenv("LOCALAPPDATA"), "Go")
			// 将GOROOT和%GOROOT/bin%写入环境变量
			err := tool.SetRegistryValue(ENVGOROOT, GOROOT, 1)
			if err != nil {
				tool.L.Error(err.Error())
				return err
			}
			PATH, err := tool.GetRegistryValue("PATH")
			if err != nil {
				tool.L.Error(err.Error())
				return err
			}
			// 写入环境变量
			err = tool.SetRegistryValue("PATH", "%GOROOT%\\bin;"+PATH, 2)
			if err != nil {
				tool.L.Error(err.Error())
				return err
			}
		} else { // 其他系统环境变量
			// 其他环境下指定默认的GOROOT
			GOROOT = filepath.Join(os.Getenv("HOME"), "Go")
			// 格式化好要追加的环境变量
			envStr := fmt.Sprintf(
				"\n%s\n%s\n",
				"export GOROOT=$HOME/Go",
				"export PATH=$GOROOT/bin:$PATH",
			)
			// 打开环境变量文件
			file, err := os.OpenFile(
				filepath.Join(os.Getenv("HOME"), ".bashrc"),
				os.O_CREATE|os.O_APPEND|os.O_WRONLY,
				0777,
			)
			if err != nil {
				tool.L.Error(err.Error())
				return err
			}
			defer file.Close()
			// 写入环境变量到文件中
			_, err = file.WriteString(envStr)
			if err != nil {
				tool.L.Error(err.Error())
				return err
			}
		}
		// 不管是什么操作系统，配置完环境变量后都需要刷新环境变量
		refreshEnv()
	}

	// 从环境变量中获取SDK版本列表网址BaseUrl(环境变量优先级最高)
	goSdkVerUrl := os.Getenv(ENVGOSDKVERURL)
	if strings.Contains(goSdkVerUrl, "http") {
		// 赋值获取到的环境变量
		GOSDKVERURL = goSdkVerUrl
	} else {
		// 赋值默认版本列表请求地址
		GOSDKVERURL = "https://qiniu.github.bearki.cn/gov/version/version.json"
	}

	// 从环境变量中获取SDK下载网址BaseUrl(环境变量优先级最高)
	goSdkDownUrl := os.Getenv(ENVGOSDKDOWNURL)
	if strings.Contains(goSdkDownUrl, "http") {
		// 赋值获取到的环境变量
		GOSDKDOWNURL = goSdkDownUrl
	} else {
		// 赋值SDK默认下载地址
		GOSDKDOWNURL = "https://golang.google.cn/dl/"
	}
	return nil
}
