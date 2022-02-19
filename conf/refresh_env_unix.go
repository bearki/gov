//go:build !windows

package conf

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bearki/gov/tool"
)

// refreshEnv 刷新操作系统环境变量
func refreshEnv() {
	// 拼接命令
	sourceEnv := fmt.Sprintf(
		"source %s",
		filepath.Join(os.Getenv("HOME"), ".bashrc"),
	)
	// 加载环境变量
	cmd := exec.Command(
		"/bin/bash",
		"-c",
		sourceEnv,
	)
	// 运行命令
	err := cmd.Run()
	// 接收错误信息
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	if err != nil {
		tool.L.Error("%s | %s", errBuf.String(), err.Error())
		return err
	}
}
