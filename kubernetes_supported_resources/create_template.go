package kubernetes_supported_resources

func CreateTemplate(k8sResource KubernetesResource) ([]byte, error) {
	return []byte(k8sResource.Content), nil
}
