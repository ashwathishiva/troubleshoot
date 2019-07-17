package cli

import (
	"net/url"
	"os"

	troubleshootclientv1beta1 "github.com/replicatedhq/troubleshoot/pkg/client/troubleshootclientset/typed/troubleshoot/v1beta1"
	"github.com/spf13/viper"
	"k8s.io/client-go/tools/clientcmd"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func isURL(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	return true
}

func createTroubleshootK8sClient() (*troubleshootclientv1beta1.TroubleshootV1beta1Client, error) {
	v := viper.GetViper()

	config, err := clientcmd.BuildConfigFromFlags("", v.GetString("kubecontext"))
	if err != nil {
		return nil, err
	}
	troubleshootClient, err := troubleshootclientv1beta1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return troubleshootClient, nil
}