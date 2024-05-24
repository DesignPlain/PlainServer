package iot

type Thing struct {
	// The name of the thing.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The thing type name.
	ThingTypeName string `json:"thingTypeName,omitempty" yaml:"thingTypeName,omitempty"`

	// Map of attributes of the thing.
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}
