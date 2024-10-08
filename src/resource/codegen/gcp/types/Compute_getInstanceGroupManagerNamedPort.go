package types

type Compute_getInstanceGroupManagerNamedPort struct {
	// The name of the instance group. Either `name` or `self_link` must be provided.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The port number.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
