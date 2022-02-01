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

// SetRegistryValue 设置注册表键值
// @params name    string 键名
// @params value   string 键值
// @params valType int    值类型（1-字符串值，2-可扩充字符串值）
// @return         error  错误信息
func SetRegistryValue(name string, value string, valType int) error {
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

// GetRegistryValue 获取注册表键值
// @params name  string 键名
// @return       string 键值
// @return       error  错误信息
func GetRegistryValue(name string) (string, error) {
	// 创建项目键值，若已存在，则直接返回键值，所有权限
	projectKey, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		"Environment",
		registry.ALL_ACCESS,
	)
	if err != nil {
		return "", err
	}
	// 延迟关闭
	defer projectKey.Close()
	// 从注册表获取key对应的值
	val, _, err := projectKey.GetStringValue(name)
	// 判断是否发生错误
	if err != nil {
		// 不存在时反空值
		if err == registry.ErrNotExist {
			return "", nil
		}
		return "", err
	}
	// 获取到设备SN
	return val, nil
}
