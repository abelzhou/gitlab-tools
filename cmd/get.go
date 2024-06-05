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

var GetCmd = &cobra.Command{
	Use:       "get",
	Short:     "获取信息",
	ValidArgs: []string{"project"},
	Args:      cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Long:      "诸如一些资源类的信息，譬如project等",
	Run:       getFunc,
}

func getFunc(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "project":
		projectList := gitlab.GetProject(keyword, namespace)
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
	default:
		fmt.Println(cmd.UsageString())
	}

}

func GetInit() {
	GetCmd.PersistentFlags().StringVarP(&keyword, "keyword", "k", "", "关键字")
	GetCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "命名空间")
	GetCmd.PersistentFlags().StringVarP(&split, "split", "s", "::", "分隔符")
}
