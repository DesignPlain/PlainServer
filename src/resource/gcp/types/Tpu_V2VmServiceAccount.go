package types

type Tpu_V2VmServiceAccount struct {
	// Email address of the service account. If empty, default Compute service account will be used.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	/*
	   The list of scopes to be made available for this service account. If empty, access to all
	   Cloud APIs will be allowed.
	*/
	Scopes []string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}
