package licensemanager

type LicenseGrant struct {
	// The target account for the grant in the form of the ARN for an account principal of the root user.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`

	// A list of the allowed operations for the grant. This is a subset of the allowed operations on the license.
	AllowedOperations []string `json:"allowedOperations,omitempty" yaml:"allowedOperations,omitempty"`

	// The ARN of the license to grant.
	LicenseArn string `json:"licenseArn,omitempty" yaml:"licenseArn,omitempty"`

	// The Name of the grant.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
