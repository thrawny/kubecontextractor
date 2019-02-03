package extract

import (
	"errors"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// Context extracts the given context including cluster, authinfo and namespace from the given kubeconfig
// and returns a new kubeconfig containing only the extracted data
func Context(contextName string, inConfig *clientcmdapi.Config) (*clientcmdapi.Config, error) {
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
