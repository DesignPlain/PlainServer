package verifiedaccess

import types "libds/aws/types"

type Group struct {
	// Description of the verified access group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The policy document that is associated with this resource.
	PolicyDocument string `json:"policyDocument,omitempty" yaml:"policyDocument,omitempty"`

	// Configuration block to use KMS keys for server-side encryption.
	SseConfiguration types.Verifiedaccess_GroupSseConfiguration `json:"sseConfiguration,omitempty" yaml:"sseConfiguration,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The id of the verified access instance this group is associated with.

	   The following arguments are optional:
	*/
	VerifiedaccessInstanceId string `json:"verifiedaccessInstanceId,omitempty" yaml:"verifiedaccessInstanceId,omitempty"`
}
