package types

type Appengine_FlexibleAppVersionResourcesVolume struct {
	// Unique name for the volume.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Volume size in gigabytes.
	SizeGb int `json:"sizeGb,omitempty" yaml:"sizeGb,omitempty"`

	// Underlying volume type, e.g. 'tmpfs'.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`
}
