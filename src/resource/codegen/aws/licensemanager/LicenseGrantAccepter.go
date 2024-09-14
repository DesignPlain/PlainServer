package licensemanager

type LicenseGrantAccepter struct {
	// The ARN of the grant to accept.
	GrantArn string `json:"grantArn,omitempty" yaml:"grantArn,omitempty"`
}
