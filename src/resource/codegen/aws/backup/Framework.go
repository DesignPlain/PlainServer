package backup

import types "libds/aws/types"

type Framework struct {
	// One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
	Controls []types.Backup_FrameworkControl `json:"controls,omitempty" yaml:"controls,omitempty"`

	// The description of the framework with a maximum of 1,024 characters
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Metadata that you can assign to help organize the frameworks you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
