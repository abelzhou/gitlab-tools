//
//author:abel
//date:2023/9/3
package gitlab

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"gt/pkg/http"
	"strconv"
	"strings"
	"time"
)

func GetProject(keyword string, namespace string) []Project {
	//获得全部的项目
	var retProjectList []Project
	var projectList []Project
	pageSize := 100
	page := 1
	for {

		url := fmt.Sprintf("/api/v4/projects?per_page=%d&page=%d", pageSize, page)
		err := http.GetRequest(url, &projectList)
		if err != nil {
			fmt.Println(fmt.Sprintf("ERROR 获取项目列表失败:%s", err.Error()))
		}
		if len(projectList) == 0 {
			break
		}
		page++
		for _, project := range projectList {
			if namespace != "" && project.Namespace.Name != namespace {
				continue
			}
			if strings.Contains(project.Name, keyword) {
				retProjectList = append(retProjectList, project)
			}
		}
	}

	return retProjectList
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
	listProject, resp, err := gitlabClient.Projects.ListProjects(&gitlab.ListProjectsOptions{Search: gitlab.Ptr(projectName)})
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

type Project struct {
	ID                                        int                       `json:"id"`
	Description                               string                    `json:"description"`
	Name                                      string                    `json:"name"`
	NameWithNamespace                         string                    `json:"name_with_namespace"`
	Path                                      string                    `json:"path"`
	PathWithNamespace                         string                    `json:"path_with_namespace"`
	CreatedAt                                 time.Time                 `json:"created_at"`
	DefaultBranch                             string                    `json:"default_branch"`
	TagList                                   []interface{}             `json:"tag_list"`
	Topics                                    []interface{}             `json:"topics"`
	SSHURLToRepo                              string                    `json:"ssh_url_to_repo"`
	HTTPURLToRepo                             string                    `json:"http_url_to_repo"`
	WebURL                                    string                    `json:"web_url"`
	ReadmeURL                                 string                    `json:"readme_url"`
	AvatarURL                                 interface{}               `json:"avatar_url"`
	ForksCount                                int                       `json:"forks_count"`
	StarCount                                 int                       `json:"star_count"`
	LastActivityAt                            time.Time                 `json:"last_activity_at"`
	Namespace                                 Namespace                 `json:"namespace"`
	Links                                     Links                     `json:"_links"`
	PackagesEnabled                           bool                      `json:"packages_enabled"`
	EmptyRepo                                 bool                      `json:"empty_repo"`
	Archived                                  bool                      `json:"archived"`
	Visibility                                string                    `json:"visibility"`
	ResolveOutdatedDiffDiscussions            bool                      `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy                 ContainerExpirationPolicy `json:"container_expiration_policy"`
	IssuesEnabled                             bool                      `json:"issues_enabled"`
	MergeRequestsEnabled                      bool                      `json:"merge_requests_enabled"`
	WikiEnabled                               bool                      `json:"wiki_enabled"`
	JobsEnabled                               bool                      `json:"jobs_enabled"`
	SnippetsEnabled                           bool                      `json:"snippets_enabled"`
	ContainerRegistryEnabled                  bool                      `json:"container_registry_enabled"`
	ServiceDeskEnabled                        bool                      `json:"service_desk_enabled"`
	ServiceDeskAddress                        interface{}               `json:"service_desk_address"`
	CanCreateMergeRequestIn                   bool                      `json:"can_create_merge_request_in"`
	IssuesAccessLevel                         string                    `json:"issues_access_level"`
	RepositoryAccessLevel                     string                    `json:"repository_access_level"`
	MergeRequestsAccessLevel                  string                    `json:"merge_requests_access_level"`
	ForkingAccessLevel                        string                    `json:"forking_access_level"`
	WikiAccessLevel                           string                    `json:"wiki_access_level"`
	BuildsAccessLevel                         string                    `json:"builds_access_level"`
	SnippetsAccessLevel                       string                    `json:"snippets_access_level"`
	PagesAccessLevel                          string                    `json:"pages_access_level"`
	OperationsAccessLevel                     string                    `json:"operations_access_level"`
	AnalyticsAccessLevel                      string                    `json:"analytics_access_level"`
	ContainerRegistryAccessLevel              string                    `json:"container_registry_access_level"`
	SecurityAndComplianceAccessLevel          string                    `json:"security_and_compliance_access_level"`
	EmailsDisabled                            interface{}               `json:"emails_disabled"`
	SharedRunnersEnabled                      bool                      `json:"shared_runners_enabled"`
	LfsEnabled                                bool                      `json:"lfs_enabled"`
	CreatorID                                 int                       `json:"creator_id"`
	ImportURL                                 interface{}               `json:"import_url"`
	ImportType                                interface{}               `json:"import_type"`
	ImportStatus                              string                    `json:"import_status"`
	OpenIssuesCount                           int                       `json:"open_issues_count"`
	CiDefaultGitDepth                         int                       `json:"ci_default_git_depth"`
	CiForwardDeploymentEnabled                bool                      `json:"ci_forward_deployment_enabled"`
	CiJobTokenScopeEnabled                    bool                      `json:"ci_job_token_scope_enabled"`
	PublicJobs                                bool                      `json:"public_jobs"`
	BuildTimeout                              int                       `json:"build_timeout"`
	AutoCancelPendingPipelines                string                    `json:"auto_cancel_pending_pipelines"`
	BuildCoverageRegex                        interface{}               `json:"build_coverage_regex"`
	CiConfigPath                              interface{}               `json:"ci_config_path"`
	SharedWithGroups                          []interface{}             `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool                      `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               interface{}               `json:"allow_merge_on_skipped_pipeline"`
	RestrictUserDefinedVariables              bool                      `json:"restrict_user_defined_variables"`
	RequestAccessEnabled                      bool                      `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                      `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool                      `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool                      `json:"printing_merge_request_link_enabled"`
	MergeMethod                               string                    `json:"merge_method"`
	SquashOption                              string                    `json:"squash_option"`
	SuggestionCommitMessage                   interface{}               `json:"suggestion_commit_message"`
	MergeCommitTemplate                       interface{}               `json:"merge_commit_template"`
	SquashCommitTemplate                      interface{}               `json:"squash_commit_template"`
	AutoDevopsEnabled                         bool                      `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string                    `json:"auto_devops_deploy_strategy"`
	AutocloseReferencedIssues                 bool                      `json:"autoclose_referenced_issues"`
	RepositoryStorage                         string                    `json:"repository_storage"`
	KeepLatestArtifact                        bool                      `json:"keep_latest_artifact"`
	RunnerTokenExpirationInterval             interface{}               `json:"runner_token_expiration_interval"`
	Permissions                               Permissions               `json:"permissions"`
}
type Namespace struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	Kind      string      `json:"kind"`
	FullPath  string      `json:"full_path"`
	ParentID  interface{} `json:"parent_id"`
	AvatarURL interface{} `json:"avatar_url"`
	WebURL    string      `json:"web_url"`
}
type Links struct {
	Self          string `json:"self"`
	Issues        string `json:"issues"`
	MergeRequests string `json:"merge_requests"`
	RepoBranches  string `json:"repo_branches"`
	Labels        string `json:"labels"`
	Events        string `json:"events"`
	Members       string `json:"members"`
}
type ContainerExpirationPolicy struct {
	Cadence       string      `json:"cadence"`
	Enabled       bool        `json:"enabled"`
	KeepN         int         `json:"keep_n"`
	OlderThan     string      `json:"older_than"`
	NameRegex     string      `json:"name_regex"`
	NameRegexKeep interface{} `json:"name_regex_keep"`
	NextRunAt     time.Time   `json:"next_run_at"`
}
type GroupAccess struct {
	AccessLevel       int `json:"access_level"`
	NotificationLevel int `json:"notification_level"`
}
type Permissions struct {
	ProjectAccess interface{} `json:"project_access"`
	GroupAccess   GroupAccess `json:"group_access"`
}
