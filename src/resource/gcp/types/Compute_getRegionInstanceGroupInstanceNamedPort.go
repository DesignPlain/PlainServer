package types

type Compute_getRegionInstanceGroupInstanceNamedPort struct {
	// The name of the instance group.  One of `name` or `self_link` must be provided.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Integer port number
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
