//
//author:abel
//date:2024/6/5
package gitlab

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func GetUsers(username string, printFlag bool) []*gitlab.User {
	page := 1
	orderBy := "updated_at"
	var allUsers []*gitlab.User
	if printFlag {
		fmt.Println(fmt.Sprintf("ID Name UserName State LastSignInTime"))
	}

	for {

		listUserOpt := &gitlab.ListUsersOptions{
			ListOptions: gitlab.ListOptions{PerPage: 999, Page: page},
			Search:      gitlab.Ptr(username),
			OrderBy:     &orderBy,
		}
		listUsers, resp, err := gitlabClient.Users.ListUsers(listUserOpt)
		if err != nil {
			fmt.Println(err.Error())
			return listUsers
		}
		if printFlag {
			for i := 0; i < len(listUsers); i++ {
				lastSignInTime := "Never logged in"
				if listUsers[i].LastSignInAt != nil {
					lastSignInTime = listUsers[i].LastSignInAt.Format("20060102")
				}
				fmt.Println(fmt.Sprintf("%d %s %s %s %s", listUsers[i].ID, listUsers[i].Name, listUsers[i].Username, listUsers[i].State, lastSignInTime))
			}

		}
		allUsers = append(allUsers, listUsers...)
		if page >= resp.TotalPages {
			break
		}
		page++

	}

	return allUsers
}

func GetUserProject(username, namespace string, printFlag bool, split string) []*gitlab.Project {
	var respListProject []*gitlab.Project

	// Verify username is correct
	currentUser := getOneUserByUsername(username)
	if currentUser == nil {
		return respListProject
	}
	fmt.Println(fmt.Sprintf("be found: %s %s %s", currentUser.Username, currentUser.Name, currentUser.State))

	if currentUser.State == "blocked" {
		return respListProject
	}

	// Get all projects
	listProject := GetProject("", namespace, false, "::")
	// Get all members of the project
	for _, project := range listProject {
		listUser, resp, err := gitlabClient.Projects.ListProjectsUsers(project.ID, &gitlab.ListProjectUserOptions{ListOptions: gitlab.ListOptions{PerPage: 9999}})
		if err != nil {
			fmt.Println(err.Error())
		}
		if resp.TotalItems == 0 {
			continue
		}
		for _, user := range listUser {
			// User not in project
			if user.Username != currentUser.Username {
				continue
			}
			// User status is blocked
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

// Get user by exact match
func getOneUserByUsername(username string) *gitlab.User {
	listUsers := GetUsers(username, false)
	var currentUser *gitlab.User
	for i := 0; i < len(listUsers); i++ {
		if listUsers[i].Username == username {
			currentUser = listUsers[i]
		}
	}
	if currentUser == nil {
		fmt.Println(fmt.Sprintf("User not found: %s, possible matches below:", username))
		for i := 0; i < len(listUsers); i++ {
			fmt.Println(fmt.Sprintf("%s", listUsers[i].Username))
		}
		return nil
	}
	return currentUser
}
