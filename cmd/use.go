package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bearki/gov/conf"
	"github.com/bearki/gov/tool"
	"github.com/spf13/cobra"
)

// use version command
var useCmd = &cobra.Command{
	Use:     "use [version]",
	Aliases: []string{"u"},
	Short:   "Use the specified version",
	Long: fmt.Sprintf(
		"%s\r\n%s%s",
		"Use the specified version as the development environment of the current host.",
		"If the version is not cached locally,",
		"\"go install [version]\" will be called first.",
	),
	Example: fmt.Sprintf(
		"  %s\r\n  %s\r\n  %s\r\n  %s",
		"gov use 1.17.6",
		"gov u 1.17.6",
		"gov use 1.18beta1",
		"gov u 1.18beta1",
	),
	Run: use,
}

// use version
func use(c *cobra.Command, args []string) {
	// 判断是否传入版本信息
	if len(args) == 0 {
		tool.L.Error("golang sdk version params incorrect.")
		return
	}
	// 获取对应该平台的版本信息
	version, err := getVersionList("go" + args[0])
	if err != nil {
		tool.L.Error(err.Error())
		return
	}
	// 拼接文件所在的本地路径
	goSdkFilePath := filepath.Join(conf.GOSDKPATH, "pkg", version.FileName)
	// 判断本地是否存在该压缩包
	sdkData, err := ioutil.ReadFile(goSdkFilePath)
	if err != nil {
		if !os.IsNotExist(err) { // 文件已存在，但打开失败了
			tool.L.Error(err.Error())
			return
		}
		// 文件不存在下载新版本
		goto DOWNLOADFILE
	} else if version.Sha256 != "" && tool.MathSha256(sdkData) != version.Sha256 { // 校验sha256失败
		// 文件不完整，重新下载该版本
		goto DOWNLOADFILE
	} else { // 文件已存在，并且sha256校验正确
		// 使用该版本
		goto USEVERSION
	}
	// 下载该版本
DOWNLOADFILE:
	// 下载
	install(c, args)
	// 判断是否下载成功
	if !downloadSuccess {
		return
	}
	// 使用该版本
USEVERSION:
	// 配置SDK所在目录
	sdkPath := filepath.Join(conf.GOSDKPATH, "sdk", version.Version)
	// mkdir
	err = os.MkdirAll(sdkPath, 0755)
	if err != nil {
		tool.L.Error(err.Error())
		return
	}
	// 判断是否是ZIP压缩包
	tool.L.Info("Unzipping compressed file......")
	if strings.Contains(version.FileName, "zip") {
		// 解压zip
		err = tool.DeCompressZip(goSdkFilePath, sdkPath)
		if err != nil {
			tool.L.Error(err.Error())
			return
		}
	} else {
		// 解压tar.gz，依赖于linux tar工具，后面有时间了会替换掉
		decodeCmd := exec.Command(
			"tar",
			"-zxf",
			goSdkFilePath,
			"-C",
			sdkPath,
		)
		// 重定向错误信息
		decodeErrBuf := bytes.NewBuffer(nil)
		decodeCmd.Stderr = decodeErrBuf
		err := decodeCmd.Run()
		if err != nil {
			tool.L.Error(decodeErrBuf.String())
			tool.L.Error(err.Error())
			return
		}
		// 将go文件夹中的内容直接放置在外层
		// mvCmd := exec.Command(
		// 	"mv",
		// 	filepath.Join(sdkPath, "go", "*"),
		// 	sdkPath,
		// )
		err = os.Rename(
			filepath.Join(sdkPath, "go"),
			sdkPath,
		)
		// // 重定向错误信息
		// mvErrBuf := bytes.NewBuffer(nil)
		// mvCmd.Stderr = mvErrBuf
		// err = mvCmd.Run()
		if err != nil {
			// tool.L.Error(mvErrBuf.String())
			tool.L.Error(err.Error())
			return
		}
		// 移除掉空的go文件夹
		err = os.RemoveAll(filepath.Join(sdkPath, "go"))
		if err != nil {
			tool.L.Error(err.Error())
			return
		}
	}
	// 解压完成
	tool.L.Success("The compressed file was decompressed successfully......")
	// 移除软链
	err = os.RemoveAll(conf.GOROOT)
	if err != nil {
		tool.L.Error(err.Error())
		return
	}
	// 创建软链
	var runCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		runCmd = exec.Command(
			"cmd",
			"/c",
			"mklink",
			"/J",
			conf.GOROOT,
			sdkPath,
		)
	} else {
		runCmd = exec.Command(
			"ln",
			"-bsnf",
			sdkPath,
			conf.GOROOT,
		)
	}
	// 重定向错误信息
	runErrBuf := bytes.NewBuffer(nil)
	runCmd.Stderr = runErrBuf
	err = runCmd.Run()
	if err != nil {
		tool.L.Error(runErrBuf.String())
		tool.L.Error(err.Error())
		return
	}
	// 创建成功
	tool.L.Success("Switched to version %s", args[0])
}
