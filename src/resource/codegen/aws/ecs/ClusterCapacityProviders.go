package ecs

import types "libds/aws/types"

type ClusterCapacityProviders struct {
	// Name of the ECS cluster to manage capacity providers for.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// Set of capacity provider strategies to use by default for the cluster. Detailed below.
	DefaultCapacityProviderStrategies []types.Ecs_ClusterCapacityProvidersDefaultCapacityProviderStrategy `json:"defaultCapacityProviderStrategies,omitempty" yaml:"defaultCapacityProviderStrategies,omitempty"`

	// Set of names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
	CapacityProviders []string `json:"capacityProviders,omitempty" yaml:"capacityProviders,omitempty"`
}
