//
//author:abel
//date:2024/6/5
package gitlab

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func GetUsers(username string, printFlag bool) []*gitlab.User {
	listUserOpt := &gitlab.ListUsersOptions{
		ListOptions: gitlab.ListOptions{PerPage: 999},
		Search:      gitlab.Ptr(username),
	}
	listUsers, resp, err := gitlabClient.Users.ListUsers(listUserOpt)
	if err != nil {
		fmt.Println(err.Error())
		return listUsers
	}
	if printFlag {
		for i := 0; i < len(listUsers); i++ {
			fmt.Println(fmt.Sprintf("%d %s %s ", listUsers[i].ID, listUsers[i].Name, listUsers[i].Username))
		}

		fmt.Println(fmt.Sprintf("Total: %d", resp.TotalItems))
	}

	return listUsers
}
