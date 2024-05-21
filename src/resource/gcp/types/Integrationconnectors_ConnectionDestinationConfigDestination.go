package types

type Integrationconnectors_ConnectionDestinationConfigDestination struct {
	// Host
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// port number
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Service Attachment
	ServiceAttachment string `json:"serviceAttachment,omitempty" yaml:"serviceAttachment,omitempty"`
}
