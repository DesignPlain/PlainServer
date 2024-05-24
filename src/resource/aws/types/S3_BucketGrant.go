package types

type S3_BucketGrant struct {
	// Canonical user id to grant for. Used only when `type` is `CanonicalUser`.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// List of permissions to apply for grantee. Valid values are `READ`, `WRITE`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// Type of grantee to apply for. Valid values are `CanonicalUser` and `Group`. `AmazonCustomerByEmail` is not supported.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Uri address to grant for. Used only when `type` is `Group`.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
