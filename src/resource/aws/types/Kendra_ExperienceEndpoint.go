package types

type Kendra_ExperienceEndpoint struct {
	// The endpoint of your Amazon Kendra experience.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// The type of endpoint for your Amazon Kendra experience.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`
}
