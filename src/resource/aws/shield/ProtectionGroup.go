package shield

type ProtectionGroup struct {
	// Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events.
	Aggregation string `json:"aggregation,omitempty" yaml:"aggregation,omitempty"`

	// The Amazon Resource Names (ARNs) of the resources to include in the protection group. You must set this when you set `pattern` to ARBITRARY and you must not set it for any other `pattern` setting.
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	// The criteria to use to choose the protected resources for inclusion in the group.
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`

	// The name of the protection group.
	ProtectionGroupId string `json:"protectionGroupId,omitempty" yaml:"protectionGroupId,omitempty"`

	// The resource type to include in the protection group. You must set this when you set `pattern` to BY_RESOURCE_TYPE and you must not set it for any other `pattern` setting.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
