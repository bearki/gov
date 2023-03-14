/**
 * @Title 注册表模块
 * @Desc 主要用于读写操作系统的注册表
 * @Author Bearki
 * @DateTime 2021-11-17 14:23
 */

package tool

import (
	"fmt"

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
