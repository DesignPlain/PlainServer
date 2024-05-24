package elasticache

type UserGroup struct {
	// The current supported value is `REDIS`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The ID of the user group.

	   The following arguments are optional:
	*/
	UserGroupId string `json:"userGroupId,omitempty" yaml:"userGroupId,omitempty"`

	// The list of user IDs that belong to the user group.
	UserIds []string `json:"userIds,omitempty" yaml:"userIds,omitempty"`
}
