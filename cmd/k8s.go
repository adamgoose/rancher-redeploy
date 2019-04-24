package cmd

import (
	"os"

	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getConfig(cmd *cobra.Command) *rest.Config {
	if development == false {
		c, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		return c
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
