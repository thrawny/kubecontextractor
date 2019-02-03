package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thrawny/kubectl-extract/pkg/extract"
	"k8s.io/client-go/tools/clientcmd"
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Extract a context including its cluster, authinfo and namespace",
	RunE: func(cmd *cobra.Command, args []string) error {
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		loadingRules.ExplicitPath = kubeConfig

		kubeConfig, err := loadingRules.Load()
		if err != nil {
			return err
		}

		contextName := context
		if contextName == "" {
			contextName = kubeConfig.CurrentContext
		}
		newConfig, err := extract.Context(contextName, kubeConfig)
		if err != nil {
			return err
		}
		content, err := clientcmd.Write(*newConfig)
		fmt.Print(string(content))
		return nil
	},
}
