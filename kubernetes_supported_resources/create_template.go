package kubernetes_supported_resources

import (
	"fmt"

	"github.com/nirsht/kube-sample/utils"
)

func getResourceByName(resourceName string) ([]byte, error) {
	filename := fmt.Sprintf("%v.%v", resourceName, "yaml")
	fileContent, err := utils.ReadFileFromTemplatesInGithub
	return fileContent, nil
}

func CreateTemplate(k8sResource KubernetesResource) ([]byte, error) {
	return getResourceByName(k8sResource.Name)
}
