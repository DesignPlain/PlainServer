package devicefarm

import types "libds/aws/types"

type DevicePool struct {
	// The device pool's description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The number of devices that Device Farm can add to your device pool.
	MaxDevices int `json:"maxDevices,omitempty" yaml:"maxDevices,omitempty"`

	// The name of the Device Pool
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARN of the project for the device pool.
	ProjectArn string `json:"projectArn,omitempty" yaml:"projectArn,omitempty"`

	// The device pool's rules. See Rule.
	Rules []types.Devicefarm_DevicePoolRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
