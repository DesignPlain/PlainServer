package redis

import types "libds/gcp/types"

type Cluster struct {
	// Required. Number of shards for the Redis cluster.
	ShardCount int `json:"shardCount,omitempty" yaml:"shardCount,omitempty"`

	/*
	   Optional. The in-transit encryption for the Redis cluster.
	   If not provided, encryption is disabled for the cluster.
	   Default value is `TRANSIT_ENCRYPTION_MODE_DISABLED`.
	   Possible values are: `TRANSIT_ENCRYPTION_MODE_UNSPECIFIED`, `TRANSIT_ENCRYPTION_MODE_DISABLED`, `TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION`.
	*/
	TransitEncryptionMode string `json:"transitEncryptionMode,omitempty" yaml:"transitEncryptionMode,omitempty"`

	/*
	   Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
	   Default value is `AUTH_MODE_DISABLED`.
	   Possible values are: `AUTH_MODE_UNSPECIFIED`, `AUTH_MODE_IAM_AUTH`, `AUTH_MODE_DISABLED`.
	*/
	AuthorizationMode string `json:"authorizationMode,omitempty" yaml:"authorizationMode,omitempty"`

	/*
	   Unique name of the resource in this scope including project and location using the form:
	   projects/{projectId}/locations/{locationId}/clusters/{clusterId}
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Required. Each PscConfig configures the consumer network where two
	   network addresses will be designated to the cluster for client access.
	   Currently, only one PscConfig is supported.
	   Structure is documented below.
	*/
	PscConfigs []types.Redis_ClusterPscConfig `json:"pscConfigs,omitempty" yaml:"pscConfigs,omitempty"`

	// The name of the region of the Redis cluster.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Optional. The number of replica nodes per shard.
	ReplicaCount int `json:"replicaCount,omitempty" yaml:"replicaCount,omitempty"`
}
