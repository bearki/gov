/**
 *@Title Define command
 *@Desc All commands will be defined here
 *@Author Bearki
 *@DateTime 2022/01/19 15:21
 */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Define all command
var rootCmd = &cobra.Command{}

// initial binding
func init() {
	// rootCmd Append Command
	rootCmd.AddCommand(installCmd, useCmd, removeCmd, listCmd)
}

// Run App
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
