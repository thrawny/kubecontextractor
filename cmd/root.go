package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"log"
	"os"
	"path/filepath"
)

var ContextName string
var SourceConfig string
var DestinationFilePath string

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "foo",
	Long:  "foo",
	Run: func(cmd *cobra.Command, args []string) {
		kubeConfig, err := clientcmd.LoadFromFile(SourceConfig)
		if err != nil {
			log.Fatal(err)
		}
		newConfig, err := extractContext(ContextName, kubeConfig)
		if err != nil {
			log.Fatal(err)
		}
		err = clientcmd.WriteToFile(*newConfig, DestinationFilePath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func extractContext(contextName string, inConfig *clientcmdapi.Config) (*clientcmdapi.Config, error) {
	context, ok := inConfig.Contexts[contextName]
	if !ok {
		return nil, errors.New("context does not exist")
	}
	outConfig := clientcmdapi.Config{
		Kind:           inConfig.Kind,
		APIVersion:     inConfig.APIVersion,
		Preferences:    inConfig.Preferences,
		Extensions:     inConfig.Extensions,
		CurrentContext: contextName,
		Contexts:       map[string]*clientcmdapi.Context{contextName: context},
	}
	if context.AuthInfo != "" {
		outConfig.AuthInfos = map[string]*clientcmdapi.AuthInfo{
			context.AuthInfo: inConfig.AuthInfos[context.AuthInfo],
		}
	}
	if context.Cluster != "" {
		outConfig.Clusters = map[string]*clientcmdapi.Cluster{
			context.Cluster: inConfig.Clusters[context.Cluster],
		}
	}
	return &outConfig, nil
}

func Execute() {
	defaultKubeConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	rootCmd.PersistentFlags().StringVar(&ContextName, "context-name", "", "The kube context to extract")
	rootCmd.PersistentFlags().StringVar(&SourceConfig, "kube-config", defaultKubeConfigPath, "The kube config file to extract from")
	rootCmd.PersistentFlags().StringVar(&DestinationFilePath, "destination-file", ContextName, "Path to file to write the extracted context to")
	rootCmd.MarkPersistentFlagRequired("context-name")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
