//
//author:abel
//date:2023/9/3
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gt/pkg/gitlab"
)

var (
	namespace string
	keyword   string
)

var GetCmd = &cobra.Command{
	Use:       "get",
	Short:     "获取信息",
	ValidArgs: []string{"project"},
	Args:      cobra.OnlyValidArgs,
	Long:      "诸如一些资源类的信息，譬如project,namespace,user等",
	Run:       getFunc,
}

func getFunc(cmd *cobra.Command, args []string) {
	//fmt.Printf("exec command:[%s]  keyword:[%s]  namespace:[%s]\n", args[0], keyword, namespace)
	projectList := gitlab.GetProject(keyword, namespace)
	for _, project := range projectList {
		fmt.Printf("%s::%s::%s::%s::%s::%s\n",
			project.Name,
			project.Namespace.Name,
			project.PathWithNamespace,
			project.SSHURLToRepo,
			project.CreatedAt.Format("2006-01-02 15:04:05"),
			project.Description,
		)
	}
}

func GetInit() {
	GetCmd.PersistentFlags().StringVarP(&keyword, "keyword", "k", "", "关键字")
	GetCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "命名空间")
}
