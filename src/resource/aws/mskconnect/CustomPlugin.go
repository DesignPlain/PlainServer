package mskconnect

import types "DesignSphere_Server/src/resource/aws/types"

type CustomPlugin struct {
	// The type of the plugin file. Allowed values are `ZIP` and `JAR`.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// A summary description of the custom plugin.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Information about the location of a custom plugin. See below.

	   The following arguments are optional:
	*/
	Location types.Mskconnect_CustomPluginLocation `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of the custom plugin..
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
