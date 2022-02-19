//go:build !windows

/**
 * @Title 注册表模块
 * @Desc 主要用于读写操作系统的注册表
 * @Author Bearki
 * @DateTime 2021-11-17 14:23
 */

package tool

// SetRegistryValue 设置注册表键值
// @params name    string 键名
// @params value   string 键值
// @params valType int    值类型（1-字符串值，2-可扩充字符串值）
// @return         error  错误信息
func SetRegistryValue(name string, value string, valType int) error { return nil }

// GetRegistryValue 获取注册表键值
// @params name  string 键名
// @return       string 键值
// @return       error  错误信息
func GetRegistryValue(name string) (string, error) { return "", nil }
