package rds

import types "libds/aws/types"

type ParameterGroup struct {
	// The name of the DB parameter group. If omitted, this provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The DB parameters to apply. See `parameter` Block below for more details. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-parameters.html) after initial creation of the group.
	Parameters []types.Rds_ParameterGroupParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	//
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of the DB parameter group. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The family of the DB parameter group.
	Family string `json:"family,omitempty" yaml:"family,omitempty"`
}
