package types

type Finspace_KxVolumeNas1Configuration struct {
	// The size of the network attached storage.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The type of file system volume. Currently, FinSpace only supports the `NAS_1` volume type. When you select the `NAS_1` volume type, you must also provide `nas1_configuration`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
