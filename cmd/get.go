//
//author:abel
//date:2023/9/3
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gt/pkg/gitlab"
)

var getExample = "get project -n namespace    --获取命名空间下的全部项目\n" +
	"get project {projectKeyword} -n namespace    --获取命名空间下的模糊匹配项目\n" +
	"get project {projectKeyword} -n namespace/   --递归获取命名空间下所有模糊匹配项目\n" +
	"get namespace {namespaceKeyword}    --查找命名空间\n" +
	"get user {username}    --查找用户\n" +
	"get projectuser {projectKeyword}    --获取项目的全部用户\n" +
	"get userproject {userKeyword} -n namespace   --获取用户参与的全部项目（性能较慢）\n"

var GetCmd = &cobra.Command{
	Use:       "get",
	Short:     "获取信息",
	ValidArgs: []string{"project", "namespace", "user", "projectuser", "userproject"},
	Args:      cobra.MatchAll(cobra.MinimumNArgs(1)),
	Long:      "诸如一些资源类的信息，譬如project等",
	Example:   getExample,
	Run:       getFunc,
}

func getFunc(cmd *cobra.Command, args []string) {
	keywords := ""
	if len(args) > 1 {
		keywords = args[1]
	}
	switch args[0] {
	case "project":
		//if namespace == "" {
		//	fmt.Println(cmd.UsageString())
		//	return
		//}
		gitlab.GetProject(keywords, namespace, true, split)
	case "projectuser":
		gitlab.GetProjectUser(keywords)
	case "userproject":
		gitlab.GetUserProject(keywords, namespace, true, split)
	case "namespace":
		gitlab.GetNamespace(keywords)
	case "user":
		gitlab.GetUsers(keywords, true)
	default:
		fmt.Println(cmd.UsageString())
	}

}

func GetInit() {
	GetCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "命名空间")
	GetCmd.PersistentFlags().StringVarP(&split, "split", "s", "::", "分隔符")
}
