package types

type Compute_getRegionInstanceGroupInstance struct {
	// URL to the instance.
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	// List of named ports in the group, as a list of resources, each containing:
	NamedPorts []Compute_getRegionInstanceGroupInstanceNamedPort `json:"namedPorts,omitempty" yaml:"namedPorts,omitempty"`

	// String description of current state of the instance.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
