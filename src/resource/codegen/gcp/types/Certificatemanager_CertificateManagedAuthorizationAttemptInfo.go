package types

type Certificatemanager_CertificateManagedAuthorizationAttemptInfo struct {
	/*
	   (Output)
	   Human readable explanation for reaching the state. Provided to help
	   address the configuration issues.
	   Not guaranteed to be stable. For programmatic access use `failure_reason` field.
	*/
	Details string `json:"details,omitempty" yaml:"details,omitempty"`

	/*
	   (Output)
	   Domain name of the authorization attempt.
	*/
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	/*
	   (Output)
	   Reason for failure of the authorization attempt for the domain.
	*/
	FailureReason string `json:"failureReason,omitempty" yaml:"failureReason,omitempty"`

	/*
	   (Output)
	   State of the domain for managed certificate issuance.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
