package types

type Integrationconnectors_ConnectionEventingConfigRegistrationDestinationConfigDestination struct {
	// port number
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Service Attachment
	ServiceAttachment string `json:"serviceAttachment,omitempty" yaml:"serviceAttachment,omitempty"`

	// Host
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
}
