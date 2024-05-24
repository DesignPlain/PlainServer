package types

type Kendra_getExperienceEndpoint struct {
	// Endpoint of your Amazon Kendra Experience.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// Type of endpoint for your Amazon Kendra Experience.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`
}
