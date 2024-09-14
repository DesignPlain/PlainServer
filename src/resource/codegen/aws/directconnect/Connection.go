package directconnect

type Connection struct {
	// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	Bandwidth string `json:"bandwidth,omitempty" yaml:"bandwidth,omitempty"`

	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `no_encrypt`, `should_encrypt`, and `must_encrypt`.
	EncryptionMode string `json:"encryptionMode,omitempty" yaml:"encryptionMode,omitempty"`

	// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of the connection.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name of the service provider associated with the connection.
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	/*
	   Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.

	   > --NOTE:-- Changing the value of `request_macsec` will cause the resource to be destroyed and re-created.
	*/
	RequestMacsec bool `json:"requestMacsec,omitempty" yaml:"requestMacsec,omitempty"`
}
