package types

type Appstream_ImageBuilderAccessEndpoint struct {
	// Identifier (ID) of the VPC in which the interface endpoint is used.
	VpceId string `json:"vpceId,omitempty" yaml:"vpceId,omitempty"`

	// Type of interface endpoint.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`
}
