//
//author:abel
//date:2024/6/5
package gitlab

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func GetNamespace(namespace string) {
	listNamespaceOpt := &gitlab.ListNamespacesOptions{
		ListOptions: gitlab.ListOptions{PerPage: 99999},
		Search:      gitlab.Ptr(namespace),
		OwnedOnly:   gitlab.Ptr(true),
	}
	namespaceList, resp, err := gitlabClient.Namespaces.ListNamespaces(listNamespaceOpt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(namespaceList); i++ {
		fmt.Println(fmt.Sprintf("%d: %s (%s)", namespaceList[i].ID, namespaceList[i].Name, namespaceList[i].WebURL))
	}
	fmt.Println(fmt.Sprintf("total: %d ", resp.TotalItems))
}
