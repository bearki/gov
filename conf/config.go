package conf

import (
	"os"
	"path/filepath"
)

var (
	// gov version
	Version string
	// default GOSDK path
	GOSDKPATH string
	// go version list file path
	GOVERSIONFILE string
	// default GOSDK download url
	GOSDKBASEURL string = "https://golang.google.cn/dl/"
)

// init env
func init() {
	// 初始化SDK默认存放路径
	exeDirPath, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	GOSDKPATH = filepath.Dir(exeDirPath)

	// 从环境变量中获取SDK存放路径
	goSdkPath := os.Getenv("GOSDKPATH")
	if len(goSdkPath) > 0 {
		GOSDKPATH = goSdkPath
	}

	GOVERSIONFILE = filepath.Join(GOSDKPATH, "version.json")

	// 从环境变量中获取SDK下载网址BaseUrl
	goSdkBaseUrl := os.Getenv("GOSDKBASEURL")
	if len(goSdkBaseUrl) > 0 {
		GOSDKBASEURL = goSdkBaseUrl
	}
}
