package types

type Compute_RouterBgpAdvertisedIpRange struct {
	// User-specified description for the IP range.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The IP range to advertise. The value must be a
	   CIDR-formatted string.
	*/
	Range string `json:"range,omitempty" yaml:"range,omitempty"`
}
