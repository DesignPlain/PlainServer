package vpclattice

import types "DesignSphere_Server/src/resource/aws/types"

type TargetGroup struct {
	// The name of the target group. The name must be unique within the account. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The type of target group. Valid Values are `IP` | `LAMBDA` | `INSTANCE` | `ALB`

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The target group configuration.
	Config types.Vpclattice_TargetGroupConfig `json:"config,omitempty" yaml:"config,omitempty"`
}
