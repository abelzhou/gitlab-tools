//
//author:abel
//date:2023/9/3
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gt/pkg/gitlab"
	"strings"
)

var getExample = "get project -n namespace \n" +
	"get project {projectKeyword} -n namespace \n" +
	"get namespace {namespaceKeyword}\n" +
	"get user {username}\n"

var GetCmd = &cobra.Command{
	Use:       "get",
	Short:     "获取信息",
	ValidArgs: []string{"project", "namespace", "user"},
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
		projectList := gitlab.GetProject(keywords, namespace)
		for _, project := range projectList {
			project.Description = strings.ReplaceAll(project.Description, "\r\n", " ")
			project.Description = strings.ReplaceAll(project.Description, "\n", " ")
			project.Description = strings.ReplaceAll(project.Description, "\r", " ")
			project.Description = strings.ReplaceAll(project.Description, "::", "：：")
			fmt.Printf("%s%s%s%s%s%s%s%s%s%s%s\n",
				project.Name,
				split,
				project.Namespace.Name,
				split,
				project.PathWithNamespace,
				split,
				project.SSHURLToRepo,
				split,
				project.CreatedAt.Format("2006-01-02 15:04:05"),
				split,
				project.Description,
			)
		}
	case "projectmember":
		gitlab.GetProjectMember(keywords)
	case "namespace":
		gitlab.GetNamespace(keywords)
	case "user":
		gitlab.GetUsers(keywords, true)
	case "userproject":
		gitlab.GetProjectByUserName(keywords)
	default:
		fmt.Println(cmd.UsageString())
	}

}

func GetInit() {
	GetCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "命名空间")
	GetCmd.PersistentFlags().StringVarP(&split, "split", "s", "::", "分隔符")
}
