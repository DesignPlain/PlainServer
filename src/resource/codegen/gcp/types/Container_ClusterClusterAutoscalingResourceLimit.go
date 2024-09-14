package types

type Container_ClusterClusterAutoscalingResourceLimit struct {
	// Minimum amount of the resource in the cluster.
	Minimum int `json:"minimum,omitempty" yaml:"minimum,omitempty"`

	/*
	   The type of the resource. For example, `cpu` and
	   `memory`.  See the [guide to using Node Auto-Provisioning](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning)
	   for a list of types.
	*/
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// Maximum amount of the resource in the cluster.
	Maximum int `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}
