package connect

type UserHierarchyGroup struct {
	/*
	   Tags to apply to the hierarchy group. If configured with a provider
	   `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// The name of the user hierarchy group. Must not be more than 100 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The identifier for the parent hierarchy group. The user hierarchy is created at level one if the parent group ID is null.
	ParentGroupId string `json:"parentGroupId,omitempty" yaml:"parentGroupId,omitempty"`
}
