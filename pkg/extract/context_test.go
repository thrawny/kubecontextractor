package extract

import (
	"github.com/stretchr/testify/assert"
	"testing"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// example config containing two contexts
func config() *clientcmdapi.Config {
	return &clientcmdapi.Config{
		Contexts: map[string]*clientcmdapi.Context{
			"foo": {
				Cluster:   "foocluster",
				AuthInfo:  "fooauthinfo",
				Namespace: "foonamespace",
			},
			"bar": {
				Cluster:   "barcluster",
				Namespace: "barnamespace",
			},
		},
		CurrentContext: "foo",
		Clusters: map[string]*clientcmdapi.Cluster{
			"foocluster": {
				Server: "http://foo.io",
			},
			"barcluster": {
				Server: "http://bar.io",
			},
		},
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			"fooauthinfo": {
				Username: "foouser",
			},
		},
	}
}

// what we expected to get if we extract the foo context
func fooConfig() *clientcmdapi.Config {
	return &clientcmdapi.Config{
		Contexts: map[string]*clientcmdapi.Context{
			"foo": {
				Cluster:   "foocluster",
				AuthInfo:  "fooauthinfo",
				Namespace: "foonamespace",
			},
		},
		CurrentContext: "foo",
		Clusters: map[string]*clientcmdapi.Cluster{
			"foocluster": {
				Server: "http://foo.io",
			},
		},
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			"fooauthinfo": {
				Username: "foouser",
			},
		},
	}
}

// what we expected to get if we extract the bar context
func barConfig() *clientcmdapi.Config {
	return &clientcmdapi.Config{
		Contexts: map[string]*clientcmdapi.Context{
			"bar": {
				Cluster:   "barcluster",
				Namespace: "barnamespace",
			},
		},
		CurrentContext: "bar",
		Clusters: map[string]*clientcmdapi.Cluster{
			"barcluster": {
				Server: "http://bar.io",
			},
		},
	}
}

func TestContext(t *testing.T) {
	type args struct {
		contextName string
		inConfig    *clientcmdapi.Config
	}

	tests := []struct {
		name      string
		args      args
		expected  *clientcmdapi.Config
		expectErr bool
	}{
		{
			name: "returns context if it exists",
			args: args{
				contextName: "foo",
				inConfig:    config(),
			},
			expected:  fooConfig(),
			expectErr: false,
		},
		{
			name: "handles missing data",
			args: args{
				contextName: "bar",
				inConfig:    config(),
			},
			expected:  barConfig(),
			expectErr: false,
		},
		{
			name: "returns error when context does not exist",
			args: args{
				contextName: "baz",
				inConfig:    config(),
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Context(tt.args.contextName, tt.args.inConfig)
			if tt.expectErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
