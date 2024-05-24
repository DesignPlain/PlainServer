package globalaccelerator

import types "DesignSphere_Server/src/resource/aws/types"

type Accelerator struct {
	// The attributes of the accelerator. Fields documented below.
	Attributes types.Globalaccelerator_AcceleratorAttributes `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// Indicates whether the accelerator is enabled. Defaults to `true`. Valid values: `true`, `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The value for the address type. Defaults to `IPV4`. Valid values: `IPV4`, `DUAL_STACK`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	// The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`

	// The name of the accelerator.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
