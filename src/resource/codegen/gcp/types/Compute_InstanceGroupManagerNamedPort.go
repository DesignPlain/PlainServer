package types

type Compute_InstanceGroupManagerNamedPort struct {
	// The name of the port.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The port number.
	   - - -
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
