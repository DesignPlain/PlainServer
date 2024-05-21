package edgenetwork

type Subnet struct {
	// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	Ipv4Cidrs []string `json:"ipv4Cidrs,omitempty" yaml:"ipv4Cidrs,omitempty"`

	// Labels associated with this resource.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	Ipv6Cidrs []string `json:"ipv6Cidrs,omitempty" yaml:"ipv6Cidrs,omitempty"`

	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the network to which this router belongs.
	   Must be of the form: `projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}`
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   A unique ID that identifies this subnet.


	   - - -
	*/
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	VlanId int `json:"vlanId,omitempty" yaml:"vlanId,omitempty"`

	// The name of the target Distributed Cloud Edge zone.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
