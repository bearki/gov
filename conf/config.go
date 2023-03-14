package conf

import (
	"path/filepath"
)

// gov version
var Version = "0.1.10"

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
	GOSDKPATH string = filepath.Join(defaultRootPath, "Gov")
	// go root path
	GOROOT string = filepath.Join(defaultRootPath, "Go")
	// go version list file path
	GOVERSIONFILE string = filepath.Join(GOSDKPATH, "version.json")
	// default GOSDK version list url
	GOSDKVERURL string = "https://qiniu.github.bearki.cn/gov/version/version.json"
	// default GOSDK download url
	GOSDKDOWNURL string = "https://golang.google.cn/dl/"
)
