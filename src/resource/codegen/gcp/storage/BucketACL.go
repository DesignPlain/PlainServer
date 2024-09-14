package storage

type BucketACL struct {
	// Configure this ACL to be the default ACL.
	DefaultAcl string `json:"defaultAcl,omitempty" yaml:"defaultAcl,omitempty"`

	// The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
	PredefinedAcl string `json:"predefinedAcl,omitempty" yaml:"predefinedAcl,omitempty"`

	// List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
	RoleEntities []string `json:"roleEntities,omitempty" yaml:"roleEntities,omitempty"`

	/*
	   The name of the bucket it applies to.

	   - - -
	*/
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
