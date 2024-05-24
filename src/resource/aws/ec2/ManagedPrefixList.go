package ec2

import types "DesignSphere_Server/src/resource/aws/types"

type ManagedPrefixList struct {
	// Address family (`IPv4` or `IPv6`) of this prefix list.
	AddressFamily string `json:"addressFamily,omitempty" yaml:"addressFamily,omitempty"`

	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries []types.Ec2_ManagedPrefixListEntry `json:"entries,omitempty" yaml:"entries,omitempty"`

	// Maximum number of entries that this prefix list can contain.
	MaxEntries int `json:"maxEntries,omitempty" yaml:"maxEntries,omitempty"`

	// Name of this resource. The name must not start with `com.amazonaws`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
