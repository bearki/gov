package conf

import (
	"os"
	"path/filepath"
)

var (
	// gov version
	Version string = "v0.0.1"
	// default GOSDK path
	GOSDKPATH string
	// go version list file path
	GOVERSIONFILE string
	// default GOSDK version list url
	GOSDKVERURL string = "https://qiniu.github.bearki.cn/gov/version/version.json"
	// default GOSDK download url
	GOSDKDOWNURL string = "https://mirrors.ustc.edu.cn/golang"
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

	// 配置json文件的存放路径
	GOVERSIONFILE = filepath.Join(GOSDKPATH, "version.json")

	// 从环境变量中获取SDK版本列表网址BaseUrl
	goSdkVerUrl := os.Getenv("GOSDKVERURL")
	if len(goSdkVerUrl) > 0 {
		GOSDKDOWNURL = goSdkVerUrl
	}

	// 从环境变量中获取SDK下载网址BaseUrl
	goSdkDownUrl := os.Getenv("GOSDKDOWNURL")
	if len(goSdkDownUrl) > 0 {
		GOSDKDOWNURL = goSdkDownUrl
	}
}
