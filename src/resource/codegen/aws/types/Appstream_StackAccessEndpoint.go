package types

type Appstream_StackAccessEndpoint struct {
	// ID of the VPC in which the interface endpoint is used.
	VpceId string `json:"vpceId,omitempty" yaml:"vpceId,omitempty"`

	/*
	   Type of the interface endpoint.
	   See the [`AccessEndpoint` AWS API documentation](https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AccessEndpoint.html) for valid values.
	*/
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`
}
