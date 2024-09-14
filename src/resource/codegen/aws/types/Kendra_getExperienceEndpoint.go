package types

type Kendra_getExperienceEndpoint struct {
	// Type of endpoint for your Amazon Kendra Experience.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`

	// Endpoint of your Amazon Kendra Experience.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
}
