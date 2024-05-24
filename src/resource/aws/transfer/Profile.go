package transfer

type Profile struct {
	// The list of certificate Ids from the imported certificate operation.
	CertificateIds []string `json:"certificateIds,omitempty" yaml:"certificateIds,omitempty"`

	// The profile type should be LOCAL or PARTNER.
	ProfileType string `json:"profileType,omitempty" yaml:"profileType,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The As2Id is the AS2 name as defined in the RFC 4130. For inbound ttransfers this is the AS2 From Header for the AS2 messages sent from the partner. For Outbound messages this is the AS2 To Header for the AS2 messages sent to the partner. his ID cannot include spaces.
	As2Id string `json:"as2Id,omitempty" yaml:"as2Id,omitempty"`
}
