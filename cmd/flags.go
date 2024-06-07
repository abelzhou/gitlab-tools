//
//author:abel
//date:2024/6/4
package cmd

import gitlab2 "github.com/xanzy/go-gitlab"

var (
	namespace string
	split     string
)

var accessLevelMap = map[string]*gitlab2.AccessLevelValue{
	"dev":   gitlab2.Ptr(gitlab2.DeveloperPermissions),
	"main":  gitlab2.Ptr(gitlab2.MaintainerPermissions),
	"owner": gitlab2.Ptr(gitlab2.OwnerPermissions)}
