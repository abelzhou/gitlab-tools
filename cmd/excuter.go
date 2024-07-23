//
//author:abel
//date:2023/9/3
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var example = getExample + createExample

func Execute() {
	GetInit()
	CreateInit()
	var rootCmd = &cobra.Command{
		Use:     "app",
		Args:    cobra.MinimumNArgs(1),
		Example: example,
	}
	rootCmd.AddCommand(GetCmd)
	rootCmd.AddCommand(CreateCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}
