package storage

type DefaultObjectAccessControl struct {
	// The name of the object, if applied to an object.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`

	/*
	   The access permission for the entity.
	   Possible values are: `OWNER`, `READER`.


	   - - -
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// The name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	/*
	   The entity holding the permission, in one of the following forms:
	   - user-{{userId}}
	   - user-{{email}} (such as "user-liz@example.com")
	   - group-{{groupId}}
	   - group-{{email}} (such as "group-example@googlegroups.com")
	   - domain-{{domain}} (such as "domain-example.com")
	   - project-team-{{projectId}}
	   - allUsers
	   - allAuthenticatedUsers
	*/
	Entity string `json:"entity,omitempty" yaml:"entity,omitempty"`
}
