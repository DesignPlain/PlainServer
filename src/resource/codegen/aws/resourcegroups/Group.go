package resourcegroups

import types "libds/aws/types"

type Group struct {
	// A configuration associates the resource group with an AWS service and specifies how the service can interact with the resources in the group. See below for details.
	Configurations []types.Resourcegroups_GroupConfiguration `json:"configurations,omitempty" yaml:"configurations,omitempty"`

	// A description of the resource group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The resource group's name. A resource group name can have a maximum of 127 characters, including letters, numbers, hyphens, dots, and underscores. The name cannot start with `AWS` or `aws`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A `resource_query` block. Resource queries are documented below.
	ResourceQuery types.Resourcegroups_GroupResourceQuery `json:"resourceQuery,omitempty" yaml:"resourceQuery,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
