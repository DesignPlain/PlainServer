package connect

type SecurityProfile struct {
	// Specifies a list of permissions assigned to the security profile.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	/*
	   Tags to apply to the Security Profile. If configured with a provider
	   `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the description of the Security Profile.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// Specifies the name of the Security Profile.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
