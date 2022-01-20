package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/bearki/gov/conf"
	"github.com/bearki/gov/tool"
	"github.com/spf13/cobra"
)

// remove version command
var removeCmd = &cobra.Command{
	Use:     "remove [version]",
	Aliases: []string{"r"},
	Short:   "Remove the specified version",
	Long:    "Remove the specified version, throw an error when the version does not exist",
	Example: fmt.Sprintf(
		"  %s\r\n  %s\r\n  %s\r\n  %s",
		"gov remove 1.17.6",
		"gov r 1.17.6",
		"gov remove 1.18beta1",
		"gov r 1.18beta1",
	),
	Run: remove,
}

// remove golang sdk
func remove(c *cobra.Command, args []string) {
	// 判断是否传入版本信息
	if len(args) == 0 {
		tool.L.Error("golang sdk version params incorrect.")
		return
	}
	// 计算文件后缀
	fileExit := "tar.gz"
	if runtime.GOOS == "windows" {
		fileExit = "zip"
	}
	// 拼接版本
	goVersion := fmt.Sprintf("go%s", args[0])
	// 拼接压缩包名称
	zipFileName := fmt.Sprintf(
		"%s.%s-%s.%s",
		goVersion,
		runtime.GOOS,
		runtime.GOARCH,
		fileExit,
	)
	// 移除压缩包
	tool.L.Info("Removing archive of specified version......")
	err := os.RemoveAll(filepath.Join(conf.GOSDKPATH, "pkg", zipFileName))
	if err != nil {
		tool.L.Error(err.Error())
	} else {
		tool.L.Success("The file directory of the specified version %s was removed successfully", args[0])
	}
	// 移除文件夹
	tool.L.Info("Removing the specified version of the file directory......")
	err = os.RemoveAll(filepath.Join(conf.GOSDKPATH, "sdk", goVersion))
	if err != nil {
		tool.L.Error(err.Error())
	} else {
		tool.L.Success("Remove the specified version %s of the compressed file successfully", args[0])
	}
}
