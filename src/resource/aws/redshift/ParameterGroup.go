package redshift

import types "DesignSphere_Server/src/resource/aws/types"

type ParameterGroup struct {
	/*
	   A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   You can read more about the parameters that Redshift supports in the [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The family of the Redshift parameter group.
	Family string `json:"family,omitempty" yaml:"family,omitempty"`

	// The name of the Redshift parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of Redshift parameters to apply.
	Parameters []types.Redshift_ParameterGroupParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
