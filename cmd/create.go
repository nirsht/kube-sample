package cmd

import (
	"fmt"

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
				cobra.CheckErr(fmt.Errorf("create needs a name for the resource"))
			}

			fmt.Println("Check create command")
		},
	}
)

func init() {
}
