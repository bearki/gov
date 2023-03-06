//go:build !windows

package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bearki/gov/tool"
)

// unix环境下的默认家目录
var defaultRootPath = os.Getenv("HOME")

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

		// 配置完环境变量后都需要刷新环境变量
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
