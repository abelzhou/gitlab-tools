//
//author:abel
//date:2024/6/4
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gt/pkg/gitlab"
)

var CreateCmd = &cobra.Command{
	Use:       "create",
	Short:     "创建一些东西",
	ValidArgs: []string{"project"},
	Args:      cobra.MatchAll(cobra.ExactArgs(3)),
	Long:      "诸如创建一些资源类的信息，譬如project等",
	Example:   "create project $projectName $projectDesc -n $namespace",
	Run:       createFunc,
}

func createFunc(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "project":
		gitlab.CreateProject(namespace, args[1], args[2])
	default:
		fmt.Println(cmd.UsageString())
	}

}

func CreateInit() {
	CreateCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "命名空间")

}
