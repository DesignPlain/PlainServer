package types

type Networkmanager_getCoreNetworkPolicyDocumentAttachmentPolicyCondition struct {
	// string value
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Valid values include: `equals`, `not-equals`, `contains`, `begins-with`.
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	// Valid values include: `account-id`, `any`, `tag-value`, `tag-exists`, `resource-id`, `region`, `attachment-type`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// string value
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
