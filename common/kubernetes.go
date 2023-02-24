package common

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubernetsClient(path string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	return client, err
}
