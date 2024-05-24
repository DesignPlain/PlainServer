package types

type Chime_VoiceConnectorOrganizationRoute struct {
	// The FQDN or IP address to contact for origination traffic.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// The designated origination route port. Defaults to `5060`.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The priority associated with the host, with 1 being the highest priority. Higher priority hosts are attempted first.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The protocol to use for the origination route. Encryption-enabled Amazon Chime Voice Connectors use TCP protocol by default.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// The weight associated with the host. If hosts are equal in priority, calls are redistributed among them based on their relative weight.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
