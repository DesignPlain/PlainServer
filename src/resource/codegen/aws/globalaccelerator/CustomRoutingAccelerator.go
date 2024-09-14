package globalaccelerator

import types "libds/aws/types"

type CustomRoutingAccelerator struct {
	// The attributes of the accelerator. Fields documented below.
	Attributes types.Globalaccelerator_CustomRoutingAcceleratorAttributes `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// Indicates whether the accelerator is enabled. Defaults to `true`. Valid values: `true`, `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The IP address type that an accelerator supports. For a custom routing accelerator, the value must be `"IPV4"`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	// The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`

	// The name of a custom routing accelerator.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
