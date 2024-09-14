package iot

import types "libds/aws/types"

type ThingType struct {
	// , Configuration block that can contain the following properties of the thing type:
	Properties types.Iot_ThingTypeProperties `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// The name of the thing type.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
