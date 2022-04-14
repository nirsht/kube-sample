package cmd

import (
	"fmt"

	"github.com/nirsht/kube-sample/utils"
	"github.com/spf13/cobra"
)

var (
	clusterCmd = &cobra.Command{
		Use:     "cluster",
		Aliases: []string{},
		Short:   "Create or delete local cluster",
		Long:    "'kube-sample cluster' will create a new local cluster, you can create cluster with kind or minikube",
		Run: func(cmd *cobra.Command, args []string) {
			delete, _ := cmd.Flags().GetBool("delete")
			clusterType, _ := cmd.Flags().GetString("cluster-type")
			clusterName, _ := cmd.Flags().GetString("cluster-name")

			output, err := initializeClutser(clusterType, clusterName, delete)
			if err != nil {
				cobra.CheckErr(err)
			} else if output != nil {
				fmt.Println(string(output))
			}
			fmt.Println("Kube-sample:\nResource has been created successfully!!!")
		},
	}
)

func initializeClutser(clusterType string, clusterName string, delete bool) ([]byte, error) {
	if delete {
		switch clusterType {
		case "kind":
			return utils.Run("kind", "delete", "cluster", "--name", clusterName)
		case "minikube":
			return utils.Run("minikube", "delete", "-p", clusterName)
		}
	}
	switch clusterType {
	case "kind":
		return utils.Run("kind", "create", "cluster", "--name", clusterName)
	case "minikube":
		return utils.Run("minikube", "start", "-p", clusterName, "--driver=docker")
	}
	return nil, nil
}
