package storage

type DefaultObjectACL struct {
	// The name of the bucket it applies to.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	/*
	   List of role/entity pairs in the form `ROLE:entity`.
	   See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	   Omitting the field is the same as providing an empty list.
	*/
	RoleEntities []string `json:"roleEntities,omitempty" yaml:"roleEntities,omitempty"`
}
