package diagflow

type CxVersion struct {
	/*
	   The Flow to create an Version for.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The human-readable name of the version. Limit of 64 characters.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
