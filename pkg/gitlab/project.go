//
//author:abel
//date:2023/9/3
package gitlab

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"strconv"
	"strings"
)

func GetProject(keyword string, namespace string) []*gitlab.Project {

	var retListProject []*gitlab.Project
	listProject, _, err := gitlabClient.Projects.ListProjects(
		&gitlab.ListProjectsOptions{
			ListOptions: gitlab.ListOptions{PerPage: 9999},
			Search:      gitlab.Ptr(keyword),
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if namespace == "" {
		return listProject
	}

	for i := 0; i < len(listProject); i++ {
		if listProject[i].Namespace.Name != namespace {
			continue
		}
		retListProject = append(retListProject, listProject[i])
	}

	return retListProject

}

func CreateProject(namespace, name, desc string) {
	name = strings.TrimSpace(name)
	desc = strings.TrimSpace(desc)
	if len(name) == 0 {
		fmt.Println("项目名称不能为空")
		return
	}
	if len(desc) == 0 {
		fmt.Println("项目描述不能为空")
		return
	}

	//查询namespace
	namespaceList, resp, err := gitlabClient.Namespaces.SearchNamespace(namespace)
	if err != nil {
		fmt.Println("查询命名空间错误")
		return
	}
	if resp.TotalItems <= 1 {
		fmt.Println("未检索到命名空间")
		return
	}

	if len(namespaceList) <= 1 {
		fmt.Println("未检索到命名空间列表")
		return
	}

	var namespaceObj *gitlab.Namespace
	for i := 0; i < len(namespaceList); i++ {
		if namespaceList[i].Name == namespace {
			namespaceObj = namespaceList[i]
			break
		}
	}

	if namespaceObj == nil {
		fmt.Println(fmt.Sprintf("未找到命名空间:%s", namespace))
	}

	p := &gitlab.CreateProjectOptions{
		Name:                 gitlab.Ptr(name),
		NamespaceID:          gitlab.Ptr(namespaceObj.ID),
		Description:          gitlab.Ptr(desc),
		InitializeWithReadme: gitlab.Ptr(true),
		Visibility:           gitlab.Ptr(gitlab.PrivateVisibility),
	}

	project, resp, err := gitlabClient.Projects.CreateProject(p)
	if err != nil {
		fmt.Println("项目创建失败")
		fmt.Println(err.Error())
		return
	}
	if project.ID != 0 {
		fmt.Println(fmt.Sprintf("项目创建成功! \nNAME: %s \nNAMESPACE: %s \nID: %d \nPATH: %s ", project.Name, project.Namespace.Name, project.ID, project.WebURL))
		return
	}

	fmt.Println("项目创建失败")

}

func GetProjectMember(projectName string) {
	currentProject := getOneProjectByName(projectName)
	if currentProject == nil {
		return
	}

	fmt.Println(fmt.Sprintf("ProjectName: %s\nNamespace: %s", currentProject.Name, currentProject.Namespace.Name))

	projectUsers, resp, err := gitlabClient.Projects.ListProjectsUsers(currentProject.ID, &gitlab.ListProjectUserOptions{ListOptions: gitlab.ListOptions{PerPage: 999}})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp.TotalItems == 0 {
		fmt.Println("该项目不存在用户")
		return
	}
	for i := 0; i < len(projectUsers); i++ {
		fmt.Println(fmt.Sprintf("%d %s %s %s", projectUsers[i].ID, projectUsers[i].Username, projectUsers[i].Username, projectUsers[i].State))
	}

}

func getOneProjectByName(projectName string) *gitlab.Project {
	listProject, resp, err := gitlabClient.Projects.ListProjects(&gitlab.ListProjectsOptions{ListOptions: gitlab.ListOptions{PerPage: 9999}, Search: gitlab.Ptr(projectName)})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if resp.TotalItems == 0 {
		fmt.Println("找不到对应的项目")
		return nil
	}
	if len(listProject) < 1 {
		fmt.Println("找不到对应的项目列表")
		return nil
	}

	var currentProject *gitlab.Project
	for i := 0; i < len(listProject); i++ {
		if listProject[i].Name == projectName {
			currentProject = listProject[i]
		}
	}
	if currentProject == nil {
		fmt.Println("没有匹配到项目，检查是否为如下项目：")
		for i := 0; i < len(listProject); i++ {
			fmt.Println(fmt.Sprintf("%d %s %s", listProject[i].ID, listProject[i].Namespace.Name, listProject[i].Name))
		}
		return nil
	}
	return currentProject
}

func AddInvites(projectName string, accessLevel *gitlab.AccessLevelValue, usernames string) {
	currentProject := getOneProjectByName(projectName)
	if currentProject == nil {
		return
	}
	listUsername := strings.Split(usernames, ",")
	var listRequiredUsername []string
	for i := 0; i < len(listUsername); i++ {
		currentUser := getOneUserByUsername(listUsername[i])
		if currentUser == nil {
			continue
		}
		listRequiredUsername = append(listRequiredUsername, strconv.Itoa(currentUser.ID))
	}

	if len(listRequiredUsername) == 0 || len(listRequiredUsername) != len(listUsername) {
		fmt.Println("添加失败")
		return
	}

	userIds := strings.Join(listRequiredUsername, ",")

	invitesOpt := &gitlab.InvitesOptions{
		UserID:      userIds,
		AccessLevel: accessLevel,
	}

	ret, _, err := gitlabClient.Invites.ProjectInvites(currentProject.ID, invitesOpt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if ret != nil {
		fmt.Println(fmt.Sprintf("%s \n %v", ret.Status, ret.Message))
	} else {
		fmt.Println("添加失败！")
	}
}

// 根据用户名获取用户下的项目 (未完成)
func GetProjectByUserName(username string) {
	currentUser := getOneUserByUsername(username)
	if currentUser == nil {
		return
	}

	projects, resp, err := gitlabClient.Projects.ListUserContributedProjects(currentUser.ID, &gitlab.ListProjectsOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < len(projects); i++ {
		fmt.Println(fmt.Sprintf("%d %s %v", projects[i].ID, projects[i].Namespace.Name, projects[i].Name))
	}
	fmt.Println(fmt.Sprintf("total: %d", resp.TotalItems))
}

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
