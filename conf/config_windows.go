//go:build windows

package conf

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bearki/gov/tool"
)

// windows环境下的默认环境
var defaultRootPath = os.Getenv("LOCALAPPDATA")

// init env
func Init() error {
	// 从环境变量中获取SDK存放路径(环境变量优先级最高)
	goSdkPath := os.Getenv(ENVGOSDKPATH)
	if len(goSdkPath) > 0 {
		// 赋值获取到的环境变量
		GOSDKPATH = goSdkPath
		// 配置版本列表json文件的存放路径
		GOVERSIONFILE = filepath.Join(GOSDKPATH, "version.json")
	}

	// 从环境变量中获取Go Root存放路径(环境变量优先级最高)
	goRootPath := os.Getenv(ENVGOROOT)
	if len(goRootPath) > 0 {
		// 赋值获取到的环境变量
		GOROOT = goRootPath
	} else {
		// 将GOROOT写入环境变量
		err := WriteEnv(ENVGOROOT, GOROOT, 1)
		if err != nil {
			tool.L.Error(err.Error())
			return err
		}
		// // 将%GOROOT/bin%写入PATH环境变量
		err = WriteEnv("PATH", "%GOROOT%\\bin;"+os.Getenv("PATH"), 2)
		if err != nil {
			tool.L.Error(err.Error())
			return err
		}
		// 不管是什么操作系统，配置完环境变量后都需要刷新环境变量
		refreshEnv()
	}

	// 从环境变量中获取SDK版本列表网址BaseUrl(环境变量优先级最高)
	goSdkVerUrl := os.Getenv(ENVGOSDKVERURL)
	if strings.HasPrefix(goSdkVerUrl, "http") {
		// 赋值获取到的环境变量
		GOSDKVERURL = goSdkVerUrl
	}

	// 从环境变量中获取SDK下载网址BaseUrl(环境变量优先级最高)
	goSdkDownUrl := os.Getenv(ENVGOSDKDOWNURL)
	if strings.HasPrefix(goSdkDownUrl, "http") {
		// 赋值获取到的环境变量
		GOSDKDOWNURL = goSdkDownUrl
	}

	// OK
	return nil
}
