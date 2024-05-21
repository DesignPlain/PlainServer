package types

type Gkeonprem_VMwareClusterValidationCheckStatusResult struct {
	/*
	   (Output)
	   The category of the validation.
	*/
	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	// A human readable description of this VMware User Cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

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
	   Machine-readable message indicating details about last transition.
	*/
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
}
