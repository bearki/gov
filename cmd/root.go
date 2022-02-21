/**
 *@Title command root application
 *@Desc 根命令将会在该文件中定义
 *@Author Bearki
 *@DateTime 2022/01/19 15:21
 */

package cmd

import (
	"bytes"

	"github.com/bearki/gov/tool"
	"github.com/spf13/cobra"
)

// Define all command
var rootCmd = &cobra.Command{}

// initial binding
func init() {
	// 隐藏默认的工具命令
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	// rootCmd Append Command
	rootCmd.AddCommand(installCmd, useCmd, removeCmd, listCmd)
}

// Run App
func Execute() {
	// 拦截错误输出
	var errBuf bytes.Buffer
	rootCmd.SetErr(&errBuf)
	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		tool.L.Error(err.Error())
	}
}

// GetCmdNameMap 获取所有自定义的命令
func GetCmdNameMap() map[string]struct{} {
	cmdMap := make(map[string]struct{})
	for _, item := range rootCmd.Commands() {
		cmdMap[item.Name()] = struct{}{}
		for _, items := range item.Aliases {
			cmdMap[items] = struct{}{}
		}
	}
	return cmdMap
}
