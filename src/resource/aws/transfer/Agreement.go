package transfer

type Agreement struct {
	// The IAM Role which provides read and write access to the parent directory of the file location mentioned in the StartFileTransfer request.
	AccessRole string `json:"accessRole,omitempty" yaml:"accessRole,omitempty"`

	// The landing directory for the files transferred by using the AS2 protocol.
	BaseDirectory string `json:"baseDirectory,omitempty" yaml:"baseDirectory,omitempty"`

	// The Optional description of the transdfer.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The unique identifier for the AS2 local profile.
	LocalProfileId string `json:"localProfileId,omitempty" yaml:"localProfileId,omitempty"`

	// The unique identifier for the AS2 partner profile.
	PartnerProfileId string `json:"partnerProfileId,omitempty" yaml:"partnerProfileId,omitempty"`

	// The unique server identifier for the server instance. This is the specific server the agreement uses.
	ServerId string `json:"serverId,omitempty" yaml:"serverId,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
