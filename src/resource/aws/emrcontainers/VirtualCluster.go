package emrcontainers

import types "DesignSphere_Server/src/resource/aws/types"

type VirtualCluster struct {
	// Configuration block for the container provider associated with your cluster.
	ContainerProvider types.Emrcontainers_VirtualClusterContainerProvider `json:"containerProvider,omitempty" yaml:"containerProvider,omitempty"`

	// Name of the virtual cluster.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
