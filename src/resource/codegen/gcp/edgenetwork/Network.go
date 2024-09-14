package edgenetwork

type Network struct {
	/*
	   A unique ID that identifies this network.


	   - - -
	*/
	NetworkId string `json:"networkId,omitempty" yaml:"networkId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of the target Distributed Cloud Edge zone.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Labels associated with this resource.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// IP (L3) MTU value of the network. Default value is `1500`. Possible values are: `1500`, `9000`.
	Mtu int `json:"mtu,omitempty" yaml:"mtu,omitempty"`
}
