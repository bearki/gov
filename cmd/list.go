package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	"github.com/bearki/gov/conf"
	"github.com/bearki/gov/tool"
	"github.com/spf13/cobra"
)

// list version command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Show all version information",
	Long: fmt.Sprintf(
		"%s\r\n%s",
		"This will display all version information, ",
		"you can download the corresponding version SDK through this information.",
	),
	Example: fmt.Sprintf(
		"  %s\r\n  %s\r\n  %s\r\n  %s",
		"gov list",
		"gov l",
		"gov list -s yes",
		"gov l --show not",
	),
	Run: list,
}

// 需要显示的内容（all | yes | not）
var show string

func init() {
	// binding install arch
	listCmd.Flags().StringVarP(
		&show,
		"show",
		"s",
		"all",
		"Control which version needs to be displayed, [all: all versions, yes: installed version, not: not installed version]",
	)
}

// Install the new SDK version
func list(c *cobra.Command, args []string) {
	// 文件名MAP
	var tempVersionMap = make(map[string]struct{})
	// 获取已安装的版本文件列表
	err := filepath.Walk(filepath.Join(conf.GOSDKPATH, "pkg"), func(path string, info fs.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if info.IsDir() {
			return nil
		}
		tempVersionMap[filepath.Base(path)] = struct{}{}
		return nil
	})
	if err != nil && !os.IsNotExist(err) {
		tool.L.Error(err.Error())
		return
	}

	// 压缩包后缀
	fileExt := "tar.gz"
	if runtime.GOOS == "windows" {
		fileExt = "zip"
	}

	// 获取远程版本列表
	versionList, err := getRemoteVersionList()
	if err != nil {
		tool.L.Error(err.Error())
		// 获取本地版本列表
		versionList, err = getLocalVersionList()
		if err != nil {
			tool.L.Error(err.Error())
			return
		}
	}
	if len(versionList) == 0 {
		tool.L.Error("The obtained version list is empty")
		return
	}

	// 判断要显示哪些列表
	switch show {

	case "yes": // 仅显示已安装
		for _, item := range versionList {
			// 拼接文件名
			filename := fmt.Sprintf(
				"%s.%s-%s.%s",
				item.Version,
				runtime.GOOS,
				runtime.GOARCH,
				fileExt,
			)
			// 判断文件是否存在
			if _, ok := tempVersionMap[filename]; ok {
				// 已安装
				tool.L.Info("%-20s %-40s [installed]", item.Version, filename)
			}
		}

	case "not": // 仅显示未安装
		for _, item := range versionList {
			// 拼接文件名
			filename := fmt.Sprintf(
				"%s.%s-%s.%s",
				item.Version,
				runtime.GOOS,
				runtime.GOARCH,
				fileExt,
			)
			// 判断文件是否存在
			if _, ok := tempVersionMap[filename]; !ok {
				// 未安装
				tool.L.Info("%-20s %-40s [not installed]", item.Version, filename)
			}
		}

	default: // 显示全部
		for _, item := range versionList {
			// 拼接文件名
			filename := fmt.Sprintf(
				"%s.%s-%s.%s",
				item.Version,
				runtime.GOOS,
				runtime.GOARCH,
				fileExt,
			)
			// 判断文件是否存在
			if _, ok := tempVersionMap[filename]; ok {
				// 已安装
				tool.L.Trace("%-20s %-40s [installed]", item.Version, filename)
			} else {
				// 未安装
				tool.L.Info("%-20s %-40s [not installed]", item.Version, filename)
			}
		}

	}
}
