package directconnect

type LinkAggregationGroup struct {
	// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	ConnectionsBandwidth string `json:"connectionsBandwidth,omitempty" yaml:"connectionsBandwidth,omitempty"`

	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are -not- recoverable.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of the LAG.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name of the service provider associated with the LAG.
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of an existing dedicated connection to migrate to the LAG.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`
}
