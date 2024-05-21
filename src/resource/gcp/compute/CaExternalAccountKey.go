package compute

type CaExternalAccountKey struct {
	// Location for the externalAccountKey. Currently only `global` is supported.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
