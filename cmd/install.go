/**
 *@Title Download golang sdk
 *@Desc If the version already exists locally,
 *		it will no longer be downloaded
 *@Author Bearki
 *@DateTime 2022/01/19 15:21
 */

package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/bearki/beclient"
	"github.com/bearki/gov/conf"
	"github.com/bearki/gov/tool"
	"github.com/spf13/cobra"
)

// install version command
var installCmd = &cobra.Command{
	Use:     "install [version]",
	Aliases: []string{"i"},
	Short:   "Download the new Golang SDK version locally",
	Long: fmt.Sprintf(
		"%s\r\n%s",
		"Pre-check whether the version exists locally.",
		"If it does not exist, download the new Golang SDK version to the local",
	),
	Example: fmt.Sprintf(
		"  %s\r\n  %s\r\n  %s\r\n  %s",
		"gov install 1.17.6",
		"gov i 1.17.6",
		"gov install 1.18beta1",
		"gov i 1.18beta1",
	),
	Run: install,
}

// 是否下载成功
var downloadSuccess = false

// Install the new SDK version
func install(c *cobra.Command, args []string) {
	// 判断是否传入了版本号
	if len(args) != 1 {
		tool.L.Error("golang sdk version params incorrect.")
		return
	}
	// 判断文件夹是否存在
	_, err := os.Stat(conf.GOSDKPATH)
	if err != nil {
		// 文件夹存在，但发生了错误
		if !os.IsNotExist(err) {
			tool.L.Error(err.Error())
			return
		}
		// 创建文件夹
		err = os.MkdirAll(conf.GOSDKPATH, 0755)
		if err != nil {
			tool.L.Error(err.Error())
			return
		}
	}

	// 获取对应该平台的版本信息
	version, err := getVersionList("go" + args[0])
	if err != nil {
		tool.L.Error(err.Error())
		return
	}

	// 判断pkg文件夹是否存在了该版本的压缩包
	err = filepath.WalkDir(conf.GOSDKPATH, func(path string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		// 判断是否是文件夹
		if d.IsDir() {
			// 跳过
			return nil
		}
		// 判断是否包含该文件
		tempFilePath := filepath.Join(conf.GOSDKPATH, version.FileName)
		if tempFilePath == path {
			// 获取该文件的信息
			tempFile, err := os.Stat(tempFilePath)
			if err != nil {
				return err
			}
			// 判断大小是否一致
			if tempFile.Size() == version.Size {
				return fmt.Errorf("this version already exists, please use \"gov use %s\"", args[0])
			}
			// 判断sha256是否一致
			fileData, err := ioutil.ReadFile(tempFilePath)
			if err != nil {
				return err
			}
			if tool.MathSha256(fileData) == version.Sha256 {
				return fmt.Errorf("this version already exists, please use \"gov use %s\"", args[0])
			}
		}
		return nil
	})
	// 文件已存在或遍历文件夹时发生了错误
	if err != nil {
		tool.L.Error(err.Error())
		return
	}
	// 定义下载地址
	downloadUrl := fmt.Sprintf("%s/%s", conf.GOSDKDOWNURL, version.FileName)
	// 定义保存路径
	savePath := filepath.Join(conf.GOSDKPATH, "pkg", version.FileName)
	// 开始下载
	err = beclient.New(downloadUrl).
		TimeOut(time.Hour).
		Download(savePath, func(currSize, totalSize float64) {
			t := (currSize / totalSize) * 100
			fmt.Printf("%s downloading..................%.2f%%\r", version.FileName, t)
		}).
		Get(nil)
	// 从新的一行开始
	fmt.Println()
	if err != nil {
		// 下载失败了
		tool.L.Error(fmt.Sprintf("Download Error\r\n%s", err.Error()))
		return
	}
	// 判断hash256是否一致
	fileData, err := ioutil.ReadFile(savePath)
	if err != nil {
		tool.L.Error(err.Error())
		return
	}
	if version.Sha256 != tool.MathSha256(fileData) {
		tool.L.Error("The downloaded file is incomplete")
		return
	}
	// 下载成功了
	tool.L.Success("%s download success", version.FileName)
	downloadSuccess = true
}
