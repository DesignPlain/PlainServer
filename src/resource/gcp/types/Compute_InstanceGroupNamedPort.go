package types

type Compute_InstanceGroupNamedPort struct {
	// The name which the port will be mapped to.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The port number to map the name to.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
