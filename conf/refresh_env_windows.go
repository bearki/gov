//go:build windows

package conf

import (
	"syscall"
	"unsafe"

	"github.com/bearki/gov/tool"
)

// refreshEnv 刷新操作系统环境变量
func refreshEnv() {
	// 借助user32.dll库的函数刷新环境变量
	user32, err := syscall.LoadDLL("user32.dll")
	if err != nil {
		tool.L.Error(err.Error())
		return
	}
	// 加载通知发送函数
	SendMessageTimeout, err := user32.FindProc("SendMessageW")
	if err != nil {
		tool.L.Error(err.Error())
		return
	}
	// 定义微软定义好的操作码
	var HWND_BROADCAST = 0xffff
	var WM_SETTINGCHANGE = 0x001A
	// 附加消息信息，转码一下
	content, _ := syscall.UTF16FromString("Environment")
	// 调用函数，通知全局刷新环境变量
	_, _, _ = SendMessageTimeout.Call(
		uintptr(HWND_BROADCAST),
		uintptr(WM_SETTINGCHANGE),
		uintptr(0),
		uintptr(unsafe.Pointer(&content[0])),
	)
}
