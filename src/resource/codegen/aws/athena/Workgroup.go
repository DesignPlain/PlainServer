package athena

import types "libds/aws/types"

type Workgroup struct {
	// Key-value map of resource tags for the workgroup. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block with various settings for the workgroup. Documented below.
	Configuration types.Athena_WorkgroupConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Description of the workgroup.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// Name of the workgroup.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
