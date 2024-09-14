package types

type Appstream_ImageBuilderAccessEndpoint struct {
	// Type of interface endpoint. For valid values, refer to the [AWS documentation](https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AccessEndpoint.html).
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`

	// Identifier (ID) of the interface VPC endpoint.
	VpceId string `json:"vpceId,omitempty" yaml:"vpceId,omitempty"`
}
