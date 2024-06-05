//
//author:abel
//date:2024/6/4
package gitlab

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"gt/pkg/consts"
)

var gitlabClient *gitlab.Client

func InitGitlabClient() {
	gitlabClient = getClient()
}

func getClient() *gitlab.Client {
	token := viper.GetString(consts.GITLIB_TOKEN)
	url := viper.GetString(consts.GITLAB_URL)
	baseURL := fmt.Sprintf("%s/api/v4", url)
	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	if err != nil {
		panic(err)
	}
	return client
}
