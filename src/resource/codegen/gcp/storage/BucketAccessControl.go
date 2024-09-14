package storage

type BucketAccessControl struct {
	/*
	   The entity holding the permission, in one of the following forms:
	   user-userId
	   user-email
	   group-groupId
	   group-email
	   domain-domain
	   project-team-projectId
	   allUsers
	   allAuthenticatedUsers
	   Examples:
	   The user liz@example.com would be user-liz@example.com.
	   The group example@googlegroups.com would be
	   group-example@googlegroups.com.
	   To refer to all members of the Google Apps for Business domain
	   example.com, the entity would be domain-example.com.


	   - - -
	*/
	Entity string `json:"entity,omitempty" yaml:"entity,omitempty"`

	/*
	   The access permission for the entity.
	   Possible values are: `OWNER`, `READER`, `WRITER`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// The name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
