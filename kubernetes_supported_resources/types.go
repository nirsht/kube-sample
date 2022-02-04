package kubernetes_supported_resources

import (
	"fmt"
	"strings"

	"github.com/nirsht/kube-sample/kubernetes_supported_resources/manifest"
	"github.com/nirsht/kube-sample/utils"
)

type KubernetesResource struct {
	Name    string
	Content string
	Aliases []string
}

const deployment = "Deployment"
const daemonSet = "DaemonSet"
const pod = "Pod"

var resourcesMap = map[string]KubernetesResource{
	deployment: {Name: deployment, Content: strings.TrimSpace(manifest.Deployment), Aliases: []string{"deploy", "deployments", strings.ToLower(deployment)}},
	daemonSet:  {Name: daemonSet, Content: strings.TrimSpace(manifest.DaemonSet), Aliases: []string{"ds", "deployments", strings.ToLower(daemonSet)}},
	pod:        {Name: pod, Content: strings.TrimSpace(manifest.Pod), Aliases: []string{"pods", strings.ToLower(pod)}},
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
