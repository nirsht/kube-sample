package kubernetes_supported_resources

import (
	"fmt"

	"github.com/nirsht/kube-sample/utils"
)

type KubernetesResource struct {
	Name    string
	Aliases []string
}

const deployment = "Deployment"

var resourcesMap = map[string]KubernetesResource{
	deployment: {Name: deployment, Aliases: []string{"deploy"}},
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
