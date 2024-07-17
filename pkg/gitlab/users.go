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

func GetUserProject(username, namespace string, printFlag bool, split string) []*gitlab.Project {
	var respListProject []*gitlab.Project

	//确认用户名是否正确
	currentUser := getOneUserByUsername(username)
	if currentUser == nil {
		return respListProject
	}
	fmt.Println(fmt.Sprintf("be found: %s %s %s", currentUser.Username, currentUser.Name, currentUser.State))

	if currentUser.State == "blocked" {
		return respListProject
	}

	//获得全部项目
	listProject := GetProject("", namespace, false, "::")
	//获得项目下的全部成员
	for _, project := range listProject {
		listUser, resp, err := gitlabClient.Projects.ListProjectsUsers(project.ID, &gitlab.ListProjectUserOptions{ListOptions: gitlab.ListOptions{PerPage: 9999}})
		if err != nil {
			fmt.Println(err.Error())
		}
		if resp.TotalItems == 0 {
			continue
		}
		for _, user := range listUser {
			//用户不在项目里
			if user.Username != currentUser.Username {
				continue
			}
			//用户状态是blocked
			if user.State == "blocked" {
				continue
			}

			respListProject = append(respListProject, project)
			if printFlag {
				printProjectInfo(project, split)
			}
		}
	}

	return respListProject
}

// 精确获得用户
func getOneUserByUsername(username string) *gitlab.User {
	listUsers := GetUsers(username, false)
	var currentUser *gitlab.User
	for i := 0; i < len(listUsers); i++ {
		if listUsers[i].Username == username {
			currentUser = listUsers[i]
		}
	}
	if currentUser == nil {
		fmt.Println(fmt.Sprintf("没找到用户: %s ,检索到的可能用户如下：", username))
		for i := 0; i < len(listUsers); i++ {
			fmt.Println(fmt.Sprintf("%s", listUsers[i].Username))
		}
		return nil
	}
	return currentUser
}
