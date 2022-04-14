package kubernetes_supported_resources

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/nirsht/kube-sample/utils"
)

type KubernetesResource struct {
	Kind    string
	Content []byte
	Aliases []string
}

const deployment = "Deployment"
const daemonSet = "DaemonSet"
const pod = "Pod"

//go:embed manifest/Deployment.yaml
var deploymentYaml []byte

//go:embed manifest/DaemonSet.yaml
var daemonsetYaml []byte

//go:embed manifest/Pod.yaml
var podYaml []byte

// TODO: Replace alias with aliases from kubectl api-resources
var resourcesMap = map[string]KubernetesResource{
	deployment: {deployment, deploymentYaml, []string{"deploy", "deployments", strings.ToLower(deployment)}},
	daemonSet:  {daemonSet, daemonsetYaml, []string{"ds", "deployments", strings.ToLower(daemonSet)}},
	pod:        {pod, podYaml, []string{"pods", strings.ToLower(pod)}},
}

func GetTypeByAlias(alias string) (*KubernetesResource, error) {
	for _, value := range resourcesMap {
		if utils.StringInSlice(alias, value.Aliases) || alias == value.Kind {
			return &value, nil
		}
	}
	return nil, fmt.Errorf("resource %s is not supported by this CLI", alias)
}
