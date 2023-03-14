//go:build windows

package conf

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/bearki/gov/tool"
	"golang.org/x/sys/windows/registry"
)

// WriteEnv 写入到环境变量
//
// @params name 环境变量值
//
// @params value 环境变量值
//
// @params valType 环境变量的值类型（1-字符串值，2-可扩充字符串值）
//
// @return 异常信息
func WriteEnv(name, value string, valType int) error {
	// 创建项目键值，若已存在，则直接返回键值，所有权限
	projectKey, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		"Environment",
		registry.ALL_ACCESS,
	)
	if err != nil {
		return err
	}
	defer projectKey.Close()
	// 写入key-val到注册表
	if valType == 1 {
		return projectKey.SetStringValue(name, value)
	}
	if valType == 2 {
		return projectKey.SetExpandStringValue(name, value)
	}
	return fmt.Errorf("not support")
}

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
