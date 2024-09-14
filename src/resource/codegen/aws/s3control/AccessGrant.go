package s3control

import types "libds/aws/types"

type AccessGrant struct {
	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// See Location Configuration below for more details.
	AccessGrantsLocationConfiguration types.S3control_AccessGrantAccessGrantsLocationConfiguration `json:"accessGrantsLocationConfiguration,omitempty" yaml:"accessGrantsLocationConfiguration,omitempty"`

	// The ID of the S3 Access Grants location to with the access grant is giving access.
	AccessGrantsLocationId string `json:"accessGrantsLocationId,omitempty" yaml:"accessGrantsLocationId,omitempty"`

	//
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// See Grantee below for more details.
	Grantee types.S3control_AccessGrantGrantee `json:"grantee,omitempty" yaml:"grantee,omitempty"`

	// The access grant's level of access. Valid values: `READ`, `WRITE`, `READWRITE`.
	Permission string `json:"permission,omitempty" yaml:"permission,omitempty"`

	// If you are creating an access grant that grants access to only one object, set this to `Object`. Valid values: `Object`.
	S3PrefixType string `json:"s3PrefixType,omitempty" yaml:"s3PrefixType,omitempty"`
}
