//
//author:abel
//date:2024/6/4
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	gitlab2 "github.com/xanzy/go-gitlab"
	"gt/pkg/gitlab"
	"strings"
)

var createExample = "create project {projectName} {projectDesc} -n {namespace}    --创建一个项目\n" +
	"create invites {projectName} {accessLevel[rep|dev|main|owner]} {usernames1,usernames2}    --邀请一个用户进入项目 \n"

var CreateCmd = &cobra.Command{
	Use:       "create",
	Short:     "创建一些东西",
	ValidArgs: []string{"project"},
	Args:      cobra.MatchAll(cobra.MinimumNArgs(3)),
	Long:      "诸如创建一些资源类的信息，譬如project等",
	Example:   createExample,
	Run:       createFunc,
}

func createFunc(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "project":
		namespace = strings.TrimSpace(namespace)
		if namespace == "" {
			fmt.Println(cmd.UsageString())
			return
		}
		gitlab.CreateProject(namespace, args[1], args[2])
	case "invites":
		if len(args) < 4 {
			fmt.Println(cmd.UsageString())
			return
		}
		projectName := strings.TrimSpace(args[1])
		accessLevel := strings.TrimSpace(args[2])
		usernames := strings.TrimSpace(args[3])
		if usernames == "" || accessLevel == "" || projectName == "" {
			fmt.Println(cmd.UsageString())
			return
		}

		var currentAccessLevel *gitlab2.AccessLevelValue
		if v, ok := accessLevelMap[accessLevel]; !ok {
			fmt.Println(cmd.UsageString())
			return
		} else {
			currentAccessLevel = v
		}

		gitlab.AddInvites(projectName, currentAccessLevel, usernames)
	default:
		fmt.Println(cmd.UsageString())
	}

}

func CreateInit() {
	CreateCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "命名空间")

}
