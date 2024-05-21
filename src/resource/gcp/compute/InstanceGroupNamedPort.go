package compute

type InstanceGroupNamedPort struct {
	/*
	   The name of the instance group.


	   - - -
	*/
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	/*
	   The name for this named port. The name must be 1-63 characters
	   long, and comply with RFC1035.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The port number, which can be a value between 1 and 65535.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The zone of the instance group.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
