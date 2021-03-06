package v1alpha1

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Provider for the networking.zephyr.solo.io/v1alpha1 Clientset from config
func ClientsetFromConfigProvider(cfg *rest.Config) (Clientset, error) {
	return NewClientsetFromConfig(cfg)
}

// Provider for the networking.zephyr.solo.io/v1alpha1 Clientset from client
func ClientsProvider(client client.Client) Clientset {
	return NewClientset(client)
}

// Provider for TrafficPolicyClient from Clientset
func TrafficPolicyClientFromClientsetProvider(clients Clientset) TrafficPolicyClient {
	return clients.TrafficPolicies()
}

// Provider for TrafficPolicyClient from Client
func TrafficPolicyClientProvider(client client.Client) TrafficPolicyClient {
	return NewTrafficPolicyClient(client)
}

type TrafficPolicyClientFactory func(client client.Client) TrafficPolicyClient

func TrafficPolicyClientFactoryProvider() TrafficPolicyClientFactory {
	return TrafficPolicyClientProvider
}

type TrafficPolicyClientFromConfigFactory func(cfg *rest.Config) (TrafficPolicyClient, error)

func TrafficPolicyClientFromConfigFactoryProvider() TrafficPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (TrafficPolicyClient, error) {
		clients, err := NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.TrafficPolicies(), nil
	}
}

// Provider for AccessControlPolicyClient from Clientset
func AccessControlPolicyClientFromClientsetProvider(clients Clientset) AccessControlPolicyClient {
	return clients.AccessControlPolicies()
}

// Provider for AccessControlPolicyClient from Client
func AccessControlPolicyClientProvider(client client.Client) AccessControlPolicyClient {
	return NewAccessControlPolicyClient(client)
}

type AccessControlPolicyClientFactory func(client client.Client) AccessControlPolicyClient

func AccessControlPolicyClientFactoryProvider() AccessControlPolicyClientFactory {
	return AccessControlPolicyClientProvider
}

type AccessControlPolicyClientFromConfigFactory func(cfg *rest.Config) (AccessControlPolicyClient, error)

func AccessControlPolicyClientFromConfigFactoryProvider() AccessControlPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (AccessControlPolicyClient, error) {
		clients, err := NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.AccessControlPolicies(), nil
	}
}

// Provider for VirtualMeshClient from Clientset
func VirtualMeshClientFromClientsetProvider(clients Clientset) VirtualMeshClient {
	return clients.VirtualMeshes()
}

// Provider for VirtualMeshClient from Client
func VirtualMeshClientProvider(client client.Client) VirtualMeshClient {
	return NewVirtualMeshClient(client)
}

type VirtualMeshClientFactory func(client client.Client) VirtualMeshClient

func VirtualMeshClientFactoryProvider() VirtualMeshClientFactory {
	return VirtualMeshClientProvider
}

type VirtualMeshClientFromConfigFactory func(cfg *rest.Config) (VirtualMeshClient, error)

func VirtualMeshClientFromConfigFactoryProvider() VirtualMeshClientFromConfigFactory {
	return func(cfg *rest.Config) (VirtualMeshClient, error) {
		clients, err := NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualMeshes(), nil
	}
}
