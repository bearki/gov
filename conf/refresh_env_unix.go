//go:build !windows

package conf

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bearki/gov/tool"
)

// refreshEnv 刷新操作系统环境变量
func refreshEnv() {
	// 加载环境变量
	cmd := exec.Command(
		os.Getenv("SHELL"),
		"-c",
		"source",
		filepath.Join(os.Getenv("HOME"), ".profile"),
	)
	// 运行命令
	err := cmd.Run()
	// 接收错误信息
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	if err != nil {
		tool.L.Error("%s | %s", errBuf.String(), err.Error())
		return
	}
}
