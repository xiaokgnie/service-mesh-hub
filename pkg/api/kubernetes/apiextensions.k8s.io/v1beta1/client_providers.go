package v1beta1

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Provider for the apiextensions.k8s.io/v1beta1 Clientset from config
func ClientsetFromConfigProvider(cfg *rest.Config) (Clientset, error) {
	return NewClientsetFromConfig(cfg)
}

// Provider for the apiextensions.k8s.io/v1beta1 Clientset from client
func ClientsProvider(client client.Client) Clientset {
	return NewClientset(client)
}

// Provider for CustomResourceDefinitionClient from Clientset
func CustomResourceDefinitionClientFromClientsetProvider(clients Clientset) CustomResourceDefinitionClient {
	return clients.CustomResourceDefinitions()
}

// Provider for CustomResourceDefinitionClient from Client
func CustomResourceDefinitionClientProvider(client client.Client) CustomResourceDefinitionClient {
	return NewCustomResourceDefinitionClient(client)
}

type CustomResourceDefinitionClientFactory func(client client.Client) CustomResourceDefinitionClient

func CustomResourceDefinitionClientFactoryProvider() CustomResourceDefinitionClientFactory {
	return CustomResourceDefinitionClientProvider
}

type CustomResourceDefinitionClientFromConfigFactory func(cfg *rest.Config) (CustomResourceDefinitionClient, error)

func CustomResourceDefinitionClientFromConfigFactoryProvider() CustomResourceDefinitionClientFromConfigFactory {
	return func(cfg *rest.Config) (CustomResourceDefinitionClient, error) {
		clients, err := NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.CustomResourceDefinitions(), nil
	}
}
