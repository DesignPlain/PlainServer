package types

type Rds_InstanceListenerEndpoint struct {
	// Specifies the DNS address of the DB instance.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`

	// The port on which the DB accepts connections.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
