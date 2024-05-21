package types

type Gkeonprem_BareMetalAdminClusterValidationCheckStatusResult struct {
	/*
	   (Output)
	   Detailed failure information, which might be unformatted.
	*/
	Details string `json:"details,omitempty" yaml:"details,omitempty"`

	/*
	   (Output)
	   Options used for the validation check.
	*/
	Options string `json:"options,omitempty" yaml:"options,omitempty"`

	/*
	   (Output)
	   A human-readable message of the check failure.
	*/
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	/*
	   (Output)
	   The category of the validation.
	*/
	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	// A human readable description of this Bare Metal Admin Cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
