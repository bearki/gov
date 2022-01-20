package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/bearki/beclient"
	"github.com/bearki/gov/conf"
	"github.com/bearki/gov/tool"
)

// Version go官网返回的版本信息
type Version struct {
	Version string         // 版本号
	Stable  bool           // 是否为稳定版本
	Files   []*VersionFile // 版本对应的文件列表
}

// VersionFile 版本对应的文件信息
type VersionFile struct {
	FileName string // 文件名称
	OS       string // 操作系统
	Arch     string // 系统架构
	Version  string // 版本号
	Sha256   string // sha256
	Size     int64  // 文件大小
	Kind     string // 发布类型（archive压缩包）
}

// 获取成功的全局版本信息
var getSuccessVersionFile *VersionFile

// getVersionList 从Go官网获取版本列表，并返回对应版本的文件信息
// @params versionCode     string       go版本号，例如“go1.17.6”
// @return                 *VersionFile 版本文件信息
// @return                 error        错误信息
func getVersionList(versionCode string) (*VersionFile, error) {
	// 是否已经缓存了版本信息
	if getSuccessVersionFile != nil {
		return getSuccessVersionFile, nil
	}
	// 使用已从远程拉取了版本列表
	isGetRemoteVersionList := false
	// 预定义
	var err error
	// 获取本地版本列表
	versionList, _ := getLocalVersionList()
	if len(versionList) == 0 {
		// 获取远程版本列表
		versionList, err = getRemoteVersionList()
		if err != nil {
			return nil, err
		}
		isGetRemoteVersionList = true
	}

	// 未从远程拉取时拉取一次，重新判断
REPEAT:

	// 是否存在该版本
	var version *Version = nil
	for _, v := range versionList {
		if versionCode == v.Version {
			version = v
		}
	}
	if version == nil {
		// 是否已从远程拉取了版本列表
		if isGetRemoteVersionList {
			return nil, fmt.Errorf("%s version not exist", versionCode)
		}
		// 获取远程版本列表
		versionList, err = getRemoteVersionList()
		if err != nil {
			return nil, err
		}
		isGetRemoteVersionList = true
		// 再筛选一次
		goto REPEAT
	}

	// 是否存在对应当前平台的文件
	for _, v := range version.Files {
		if v.Arch == runtime.GOARCH &&
			v.OS == runtime.GOOS &&
			v.Kind == "archive" {
			if !version.Stable {
				tool.L.Warn("This is not a stable release, please use with caution.")
			} else {
				tool.L.Trace("This is a stable release, please use it with confidence.")
			}
			getSuccessVersionFile = v
			return getSuccessVersionFile, nil
		}
	}

	// 不存在这个版本
	return nil, fmt.Errorf("the SDK for the current platform was not found in version %s", versionCode)
}

// 获取本地版本列表
func getLocalVersionList() ([]*Version, error) {
	tool.L.Info("Fetching version list from local......")
	tool.L.Info("version file: %s", conf.GOVERSIONFILE)
	// 打开文件
	data, err := ioutil.ReadFile(conf.GOVERSIONFILE)
	if err != nil {
		tool.L.Warn(err.Error())
		return nil, err
	}
	// 解析json
	var jsonData []*Version
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		tool.L.Warn(err.Error())
		return nil, err
	}
	tool.L.Success("Get version list from local successfully.")
	return jsonData, nil
}

// 获取远程的版本列表
func getRemoteVersionList() ([]*Version, error) {
	tool.L.Info("Fetching version list from remote......")
	// 获取版本信息
	var response []*Version
	err := beclient.New(conf.GOSDKVERURL).
		Query("mode", "json").
		Query("include", "all").
		ContentType(beclient.ContentTypeFormURL).
		Get(&response, beclient.ContentTypeJson)
	if err != nil {
		return nil, err
	}

	// 判断是否获取到版本列表
	if len(response) == 0 {
		return nil, fmt.Errorf("version list is null")
	}

	// 将版本列表写入到本地文件
	verData, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(conf.GOSDKPATH, 0755)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(
		conf.GOVERSIONFILE,
		verData,
		0666,
	)
	if err != nil {
		return nil, err
	}
	tool.L.Success("Get version list from remote successfully.")
	// 返回版本列表
	return response, nil
}
