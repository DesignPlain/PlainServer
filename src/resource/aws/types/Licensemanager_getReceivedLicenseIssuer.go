package types

type Licensemanager_getReceivedLicenseIssuer struct {
	// Issuer key fingerprint.
	KeyFingerprint string `json:"keyFingerprint,omitempty" yaml:"keyFingerprint,omitempty"`

	// The key name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Asymmetric KMS key from AWS Key Management Service. The KMS key must have a key usage of sign and verify, and support the RSASSA-PSS SHA-256 signing algorithm.
	SignKey string `json:"signKey,omitempty" yaml:"signKey,omitempty"`
}
