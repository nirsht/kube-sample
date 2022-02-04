package cmd

import (
	"fmt"
	"os"

	"github.com/nirsht/kube-sample/kubernetes_supported_resources"
	"github.com/nirsht/kube-sample/utils"
	"github.com/spf13/cobra"
)

var (
	createCmd = &cobra.Command{
		Use:     "create [resource name]",
		Aliases: []string{},
		Short:   "Create sample yaml file",
		Long: `Create (kube-sample create) will create a new yaml file, with a sample yaml
of you requred resource`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cobra.CheckErr(fmt.Errorf("create command needs a name of the resource"))
			}
			fileName, err := initializeCreate(args)
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Printf("Your file %q that contain resource type %q is ready", string(fileName), args[0])
		},
	}
)

func initializeCreate(args []string) ([]byte, error) {
	kubernetesResource, err := kubernetes_supported_resources.GetTypeByAlias(args[0])
	if err != nil {
		return nil, err
	}

	template, err := kubernetes_supported_resources.CreateTemplate(*kubernetesResource)
	if err != nil {
		return nil, err
	}

	targetFileName, err := utils.GenerateFileName(kubernetesResource.Name)
	if err != nil {
		return nil, err
	}

	targetFilePath := fmt.Sprintf("%v.%v", string(targetFileName), "yaml")
	if err = os.WriteFile(targetFilePath, template, 0644); err != nil {
		return nil, err
	}

	return targetFileName, err
}
