package types

type S3_BucketAclV2AccessControlPolicyGrantGrantee struct {
	// Email address of the grantee. See [Regions and Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for supported AWS regions where this argument can be specified.
	EmailAddress string `json:"emailAddress,omitempty" yaml:"emailAddress,omitempty"`

	// Canonical user ID of the grantee.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Type of grantee. Valid values: `CanonicalUser`, `AmazonCustomerByEmail`, `Group`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// URI of the grantee group.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	// Display name of the owner.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
