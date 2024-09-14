package types

type S3_BucketAclV2AccessControlPolicyOwner struct {
	// Display name of the owner.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// ID of the owner.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
