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
	cmd := exec.Command("bash", "-c", "source "+filepath.Join(os.Getenv("HOME"), ".profile"))
	// 接收错误信息
	errBuf := bytes.NewBuffer(nil)
	cmd.Stderr = errBuf
	// 运行命令
	err := cmd.Run()
	if err != nil {
		tool.L.Error("%s | %s", errBuf.String(), err.Error())
		return
	}
}
