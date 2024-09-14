package types

type Codepipeline_CustomActionTypeConfigurationProperty struct {
	// The description of the action configuration property.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether the configuration property is a key.
	Key bool `json:"key,omitempty" yaml:"key,omitempty"`

	// The name of the action configuration property.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Indicates that the property will be used in conjunction with PollForJobs.
	Queryable bool `json:"queryable,omitempty" yaml:"queryable,omitempty"`

	// Whether the configuration property is a required value.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	// Whether the configuration property is secret.
	Secret bool `json:"secret,omitempty" yaml:"secret,omitempty"`

	// The type of the configuration property. Valid values: `String`, `Number`, `Boolean`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
