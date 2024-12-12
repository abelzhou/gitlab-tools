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

func GetProject(keyword string, namespace string, printFlag bool, split string) []*gitlab.Project {

	var retListProject []*gitlab.Project
	var allListProject []*gitlab.Project
	page := 1

	for {
		listProject, resp, err := gitlabClient.Projects.ListProjects(
			&gitlab.ListProjectsOptions{
				ListOptions: gitlab.ListOptions{PerPage: 100, Page: page},
				Search:      gitlab.Ptr(keyword),
			},
		)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		allListProject = append(allListProject, listProject...)
		if page >= resp.TotalPages {
			break
		}
		page++

	}

	isPrefix := strings.HasSuffix(namespace, "/")
	if isPrefix {
		namespace = namespace[:len(namespace)-1]
	}

	for i := 0; i < len(allListProject); i++ {

		namespacePathList := strings.Split(allListProject[i].Namespace.FullPath, "/")
		if namespace != "" {
			if isPrefix && namespacePathList[0] != namespace {
				continue
			}
			if !isPrefix && allListProject[i].Namespace.Name != namespace {
				continue
			}
		}
		retListProject = append(retListProject, allListProject[i])
		if printFlag {
			printProjectInfo(allListProject[i], split)
		}
	}

	return retListProject

}

func CreateProject(namespace, name, desc string) {
	name = strings.TrimSpace(name)
	desc = strings.TrimSpace(desc)
	if len(name) == 0 {
		fmt.Println("Project name cannot be empty")
		return
	}
	if len(desc) == 0 {
		fmt.Println("Project description cannot be empty")
		return
	}

	// Query namespace
	namespaceList, resp, err := gitlabClient.Namespaces.SearchNamespace(namespace)
	if err != nil {
		fmt.Println("Error querying namespace")
		return
	}
	if resp.TotalItems <= 1 {
		fmt.Println("No namespace found")
		return
	}

	if len(namespaceList) <= 1 {
		fmt.Println("No namespace list found")
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
		fmt.Println(fmt.Sprintf("Namespace not found: %s", namespace))
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
		fmt.Println("Project creation failed")
		fmt.Println(err.Error())
		return
	}
	if project.ID != 0 {
		fmt.Println(fmt.Sprintf("Project created successfully! \nNAME: %s \nNAMESPACE: %s \nID: %d \nPATH: %s ", project.Name, project.Namespace.Name, project.ID, project.WebURL))
		return
	}

	fmt.Println("Project creation failed")

}

func GetProjectUser(projectName string) {
	currentProject := getOneProjectByName(projectName)
	if currentProject == nil {
		return
	}

	fmt.Println(fmt.Sprintf("ProjectName: %s\nNamespace: %s", currentProject.Name, currentProject.Namespace.Name))

	projectUsers, resp, err := gitlabClient.Projects.ListProjectsUsers(currentProject.ID, &gitlab.ListProjectUserOptions{ListOptions: gitlab.ListOptions{PerPage: 9999}})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp.TotalItems == 0 {
		fmt.Println("No users exist in this project")
		return
	}
	for i := 0; i < len(projectUsers); i++ {
		fmt.Println(fmt.Sprintf("%d %s %s %s", projectUsers[i].ID, projectUsers[i].Name, projectUsers[i].Username, projectUsers[i].State))
	}

}

func getOneProjectByName(projectName string) *gitlab.Project {
	listProject, resp, err := gitlabClient.Projects.ListProjects(&gitlab.ListProjectsOptions{ListOptions: gitlab.ListOptions{PerPage: 9999}, Search: gitlab.Ptr(projectName)})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if resp.TotalItems == 0 {
		fmt.Println("Project not found")
		return nil
	}
	if len(listProject) < 1 {
		fmt.Println("Project list not found")
		return nil
	}

	var currentProject *gitlab.Project
	for i := 0; i < len(listProject); i++ {
		// Some projects may have leading/trailing spaces after creation
		if strings.TrimSpace(listProject[i].Name) == projectName {
			currentProject = listProject[i]
		}
	}
	if currentProject == nil {
		fmt.Println("No matching projects found, check if it's one of the following:")
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
		fmt.Println("Addition failed")
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
		fmt.Println("Addition failed!")
	}
}

func printProjectInfo(project *gitlab.Project, split string) {
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

// Get projects under a user (Not implemented)
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
