//
//author:abel
//date:2023/9/3
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gt/pkg/gitlab"
)

var getExample = "get project -n namespace    --Get all projects in namespace\n" +
	"get project {projectKeyword} -n namespace    --Get fuzzy matching projects in namespace\n" +
	"get project {projectKeyword} -n namespace/   --Recursively get all fuzzy matching projects in namespace\n" +
	"get namespace {namespaceKeyword}    --Find namespace\n" +
	"get user {username}    --Find user\n" +
	"get projectuser {projectKeyword}    --Get all users in project\n" +
	"get userproject {userKeyword} -n namespace   --Get all projects user participates in (slower performance)\n"

var GetCmd = &cobra.Command{
	Use:       "get",
	Short:     "Get information",
	ValidArgs: []string{"project", "namespace", "user", "projectuser", "userproject"},
	Args:      cobra.MatchAll(cobra.MinimumNArgs(1)),
	Long:      "Get resource information, such as projects, etc.",
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
	GetCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace")
	GetCmd.PersistentFlags().StringVarP(&split, "split", "s", "::", "Separator")
}
