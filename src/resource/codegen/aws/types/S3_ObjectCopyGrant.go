package types

type S3_ObjectCopyGrant struct {
	// URI of the grantee group. Used only when `type` is `Group`.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	// Email address of the grantee. Used only when `type` is `AmazonCustomerByEmail`.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// Canonical user ID of the grantee. Used only when `type` is `CanonicalUser`.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// List of permissions to grant to grantee. Valid values are `READ`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	/*
	   Type of grantee. Valid values are `CanonicalUser`, `Group`, and `AmazonCustomerByEmail`.

	   This configuration block has the following optional arguments (one of the three is required):
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
