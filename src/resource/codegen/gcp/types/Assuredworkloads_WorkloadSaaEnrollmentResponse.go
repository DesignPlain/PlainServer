package types

type Assuredworkloads_WorkloadSaaEnrollmentResponse struct {
	// Indicates SAA enrollment setup error if any.
	SetupErrors []string `json:"setupErrors,omitempty" yaml:"setupErrors,omitempty"`

	// Indicates SAA enrollment status of a given workload. Possible values: SETUP_STATE_UNSPECIFIED, STATUS_PENDING, STATUS_COMPLETE
	SetupStatus string `json:"setupStatus,omitempty" yaml:"setupStatus,omitempty"`
}
