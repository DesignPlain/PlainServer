package types

type Compute_getAddressesAddress struct {
	// The IP address name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Region that should be considered to search addresses.
	   All regions are considered if missing.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The URI of the created resource.
	SelfLink string `json:"selfLink,omitempty" yaml:"selfLink,omitempty"`

	// Indicates if the address is used. Possible values are: RESERVED or IN_USE.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The IP address (for example `1.2.3.4`).
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// The IP address type, can be `EXTERNAL` or `INTERNAL`.
	AddressType string `json:"addressType,omitempty" yaml:"addressType,omitempty"`

	// The IP address description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A map containing IP labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
