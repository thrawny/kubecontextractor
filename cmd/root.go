package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var context string
var kubeConfig string

var rootCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract extracts objects like contexts from a kubeconfig file",
}

func init() {
	flags := genericclioptions.ConfigFlags{KubeConfig: &kubeConfig, Context: &context}
	flags.AddFlags(rootCmd.PersistentFlags())

	rootCmd.AddCommand(contextCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
