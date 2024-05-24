package ecs

import types "DesignSphere_Server/src/resource/aws/types"

type Cluster struct {
	// The execute command configuration for the cluster. Detailed below.
	Configuration types.Ecs_ClusterConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configures a default Service Connect namespace. Detailed below.
	ServiceConnectDefaults types.Ecs_ClusterServiceConnectDefaults `json:"serviceConnectDefaults,omitempty" yaml:"serviceConnectDefaults,omitempty"`

	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
	Settings []types.Ecs_ClusterSetting `json:"settings,omitempty" yaml:"settings,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
