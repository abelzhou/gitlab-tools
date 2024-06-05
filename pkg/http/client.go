//
//author:abel
//date:2023/9/3
package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"gt/pkg/consts"
	"net/http"
	"time"
)

var authorization = "Bearer"

var client *resty.Client

func init() {
	client = resty.New()
	client.RetryCount = 3
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))
	client.SetTimeout(time.Second * 10)
	client.SetJSONEscapeHTML(false)
}

// GetRequest 发送get请求
func GetRequest(url string, resp interface{}) error {
	token := viper.GetString(consts.GITLIB_TOKEN)
	authorization = "Bearer " + token
	url = viper.GetString(consts.GITLAB_URL) + url
	send, err := client.R().SetHeader("Authorization", authorization).SetResult(&resp).Get(url)
	if err != nil {
		return err
	}
	if send.StatusCode() != http.StatusOK {
		return err
	}
	return nil
}
