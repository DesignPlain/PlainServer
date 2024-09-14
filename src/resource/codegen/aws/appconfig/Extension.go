package appconfig

import types "libds/aws/types"

type Extension struct {
	// The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
	Parameters []types.Appconfig_ExtensionParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The action points defined in the extension. Detailed below.
	ActionPoints []types.Appconfig_ExtensionActionPoint `json:"actionPoints,omitempty" yaml:"actionPoints,omitempty"`

	// Information about the extension.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
