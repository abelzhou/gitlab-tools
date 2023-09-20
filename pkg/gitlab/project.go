//
//author:abel
//date:2023/9/3
package gitlab

import (
	"fmt"
	"gt/pkg/http"
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
