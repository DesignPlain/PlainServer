package storage

type ObjectACL struct {
	// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
	PredefinedAcl string `json:"predefinedAcl,omitempty" yaml:"predefinedAcl,omitempty"`

	/*
	   List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	   Must be set if `predefined_acl` is not.
	*/
	RoleEntities []string `json:"roleEntities,omitempty" yaml:"roleEntities,omitempty"`

	// The name of the bucket the object is stored in.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	/*
	   The name of the object to apply the acl to.

	   - - -
	*/
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
}
