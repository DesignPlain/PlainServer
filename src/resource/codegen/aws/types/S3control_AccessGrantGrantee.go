package types

type S3control_AccessGrantGrantee struct {
	// Grantee identifier.
	GranteeIdentifier string `json:"granteeIdentifier,omitempty" yaml:"granteeIdentifier,omitempty"`

	// Grantee types. Valid values: `DIRECTORY_USER`, `DIRECTORY_GROUP`, `IAM`.
	GranteeType string `json:"granteeType,omitempty" yaml:"granteeType,omitempty"`
}
