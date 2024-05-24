package kms

type CustomKeyStore struct {
	// Unique name for Custom Key Store.
	CustomKeyStoreName string `json:"customKeyStoreName,omitempty" yaml:"customKeyStoreName,omitempty"`

	// Password for `kmsuser` on CloudHSM.
	KeyStorePassword string `json:"keyStorePassword,omitempty" yaml:"keyStorePassword,omitempty"`

	// Customer certificate used for signing on CloudHSM.
	TrustAnchorCertificate string `json:"trustAnchorCertificate,omitempty" yaml:"trustAnchorCertificate,omitempty"`

	// Cluster ID of CloudHSM.
	CloudHsmClusterId string `json:"cloudHsmClusterId,omitempty" yaml:"cloudHsmClusterId,omitempty"`
}
