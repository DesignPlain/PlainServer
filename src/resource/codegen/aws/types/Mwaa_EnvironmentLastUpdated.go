package types

type Mwaa_EnvironmentLastUpdated struct {
	// The status of the Amazon MWAA Environment
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The Created At date of the MWAA Environment
	CreatedAt string `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`

	//
	Errors []Mwaa_EnvironmentLastUpdatedError `json:"errors,omitempty" yaml:"errors,omitempty"`
}
