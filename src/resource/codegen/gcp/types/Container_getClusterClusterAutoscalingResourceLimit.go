package types

type Container_getClusterClusterAutoscalingResourceLimit struct {
	// Maximum amount of the resource in the cluster.
	Maximum int `json:"maximum,omitempty" yaml:"maximum,omitempty"`

	// Minimum amount of the resource in the cluster.
	Minimum int `json:"minimum,omitempty" yaml:"minimum,omitempty"`

	// The type of the resource. For example, cpu and memory. See the guide to using Node Auto-Provisioning for a list of types.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}
