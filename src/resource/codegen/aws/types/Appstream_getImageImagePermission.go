package types

type Appstream_getImageImagePermission struct {
	// indicated whether the image can be used for an image builder.
	AllowImageBuilder bool `json:"allowImageBuilder,omitempty" yaml:"allowImageBuilder,omitempty"`

	// Boolean indicating if the image can be used for a fleet.
	AllowFleet bool `json:"allowFleet,omitempty" yaml:"allowFleet,omitempty"`
}
