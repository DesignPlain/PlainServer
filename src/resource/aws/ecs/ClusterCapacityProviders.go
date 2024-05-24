package ecs

import types "DesignSphere_Server/src/resource/aws/types"

type ClusterCapacityProviders struct {
	// Set of names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
	CapacityProviders []string `json:"capacityProviders,omitempty" yaml:"capacityProviders,omitempty"`

	// Name of the ECS cluster to manage capacity providers for.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// Set of capacity provider strategies to use by default for the cluster. Detailed below.
	DefaultCapacityProviderStrategies []types.Ecs_ClusterCapacityProvidersDefaultCapacityProviderStrategy `json:"defaultCapacityProviderStrategies,omitempty" yaml:"defaultCapacityProviderStrategies,omitempty"`
}
