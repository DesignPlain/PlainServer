package mskconnect

import types "libds/aws/types"

type CustomPlugin struct {
	// The type of the plugin file. Allowed values are `ZIP` and `JAR`.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// A summary description of the custom plugin.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Information about the location of a custom plugin. See `location` Block for details.
	Location types.Mskconnect_CustomPluginLocation `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of the custom plugin..
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   The following arguments are optional:
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
