/**
 *@Title list command
 *@Desc list命令将会在该文件中定义
 *      该命令将会列出所有可用版本
 *@Author Bearki
 *@DateTime 2022/01/19 15:21
 */

package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

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

// Get Go SDK version list
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

	// 符合条件的版本列表
	var showVersionList []string
	for _, item := range versionList {
		// 文件名
		filename := ""
		for _, items := range item.Files {
			if items.Arch == runtime.GOARCH && items.OS == runtime.GOOS {
				filename = items.FileName
				break
			}
		}
		// 判断是否有匹配到文件
		if len(filename) == 0 {
			// 跳过
			continue
		}
		// 是否符合条件
		switch show {
		case "yes":
			// 判断文件是否存在
			if _, ok := tempVersionMap[filename]; ok {
				// 追加到需要渲染的列表中
				showVersionList = append(showVersionList, fmt.Sprintf("%-20s %-40s [installed]", item.Version, filename))
			}
		case "not":
			// 判断文件是否存在
			if _, ok := tempVersionMap[filename]; ok {
				// 追加到需要渲染的列表中
				showVersionList = append(showVersionList, fmt.Sprintf("%-20s %-40s [not installed]", item.Version, filename))
			}
		default:
			// 判断文件是否存在
			if _, ok := tempVersionMap[filename]; ok {
				// 已安装
				// 追加到需要渲染的列表中
				showVersionList = append(showVersionList, fmt.Sprintf("%-20s %-40s [installed]", item.Version, filename))
			} else {
				// 未安装
				// 追加到需要渲染的列表中
				showVersionList = append(showVersionList, fmt.Sprintf("%-20s %-40s [not installed]", item.Version, filename))
			}
		}
	}

	// 打印总量
	tool.L.Warn("SDK Total: %d", len(showVersionList))

	// 每页显示数据量
	const pageCap = 15
	// 遍历渲染列表，执行分页显示
	for i, v := range showVersionList {
		// 每满指定数据量就执行等待
		if (i+1)%pageCap == 0 {
			// 打印提示内容
			fmt.Print("Please press any key to view the rest (exit Q): ")
			var b string
			fmt.Scanln(&b)
			// 是否停止打印
			if b == "Q" || b == "q" {
				// 停止打印
				return
			}
		}
		// 根据不同状态打印不同颜色
		if strings.Contains(v, "[installed]") {
			// 已安装打印蓝色
			tool.L.Trace(v)
		} else {
			// 未安装打印白色
			tool.L.Info(v)
		}
	}
}
