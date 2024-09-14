package types

type Finspace_KxVolumeNas1Configuration struct {
	// The size of the network attached storage.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The type of the network attached storage.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
