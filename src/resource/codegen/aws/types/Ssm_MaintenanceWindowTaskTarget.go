package types

type Ssm_MaintenanceWindowTaskTarget struct {
	//
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The array of strings.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
