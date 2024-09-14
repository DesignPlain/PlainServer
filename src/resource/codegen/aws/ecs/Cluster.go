package ecs

import types "libds/aws/types"

type Cluster struct {
	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Execute command configuration for the cluster. See `configuration` Block for details.
	Configuration types.Ecs_ClusterConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	/*
	   Name of the cluster (up to 255 letters, numbers, hyphens, and underscores)

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Default Service Connect namespace. See `service_connect_defaults` Block for details.
	ServiceConnectDefaults types.Ecs_ClusterServiceConnectDefaults `json:"serviceConnectDefaults,omitempty" yaml:"serviceConnectDefaults,omitempty"`

	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. See `setting` Block for details.
	Settings []types.Ecs_ClusterSetting `json:"settings,omitempty" yaml:"settings,omitempty"`
}
