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
			run, _ := cmd.Flags().GetBool("run")
			apply, _ := cmd.Flags().GetBool("apply")

			output, err := initializeCreate(args, run, apply)
			if err != nil {
				cobra.CheckErr(err)
			} else if output != nil {
				fmt.Println(string(output))
			}
			fmt.Println("From kube-sample: ")
			fmt.Println("Resource has been created successfully!!")
		},
	}
)

func initializeCreate(args []string, run bool, apply bool) ([]byte, error) {
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

	if apply {
		return utils.Run("kubectl", "apply", "-f", targetFilePath)
	} else if run {
		return utils.Run("kubectl", "create", "-f", targetFilePath)
	}

	return nil, err
}
