package types

type Certificatemanager_CertificateManagedProvisioningIssue struct {
	/*
	   (Output)
	   Human readable explanation for reaching the state. Provided to help
	   address the configuration issues.
	   Not guaranteed to be stable. For programmatic access use `failure_reason` field.
	*/
	Details string `json:"details,omitempty" yaml:"details,omitempty"`

	/*
	   (Output)
	   Reason for provisioning failures.
	*/
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
}
