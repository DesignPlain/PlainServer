package rds

import types "DesignSphere_Server/src/resource/aws/types"

type OptionGroup struct {
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion string `json:"majorEngineVersion,omitempty" yaml:"majorEngineVersion,omitempty"`

	// Name of the option group. If omitted, the provider will assign a random, unique name. Must be lowercase, to match as it is stored in AWS.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Description of the option group.
	OptionGroupDescription string `json:"optionGroupDescription,omitempty" yaml:"optionGroupDescription,omitempty"`

	// The options to apply. See `option` Block below for more details.
	Options []types.Rds_OptionGroupOption `json:"options,omitempty" yaml:"options,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the name of the engine that this option group should be associated with.
	EngineName string `json:"engineName,omitempty" yaml:"engineName,omitempty"`
}
