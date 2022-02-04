package kubernetes_supported_resources

import (
	"fmt"

	"github.com/nirsht/kube-sample/kubernetes_supported_resources/manifest"
	"github.com/nirsht/kube-sample/utils"
)

type KubernetesResource struct {
	Name    string
	Content string
	Aliases []string
}

const deployment = "Deployment"

var resourcesMap = map[string]KubernetesResource{
	deployment: {Name: deployment, Content: manifest.Deployment, Aliases: []string{"deploy"}},
}

func GetTypeByAlias(alias string) (*KubernetesResource, error) {
	resource := &KubernetesResource{}
	for _, value := range resourcesMap {
		if utils.StringInSlice(alias, value.Aliases) || alias == value.Name {
			resource = &value
			return resource, nil
		}
	}
	return resource, fmt.Errorf("resource %v is not supported by this CLI", alias)
}
