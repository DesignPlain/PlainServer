package types

type Compute_InstanceFromMachineImageServiceAccount struct {
	// The service account e-mail address.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// A list of service scopes.
	Scopes []string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}
