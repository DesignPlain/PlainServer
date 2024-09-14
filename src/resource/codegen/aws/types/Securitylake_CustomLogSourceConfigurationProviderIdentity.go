package types

type Securitylake_CustomLogSourceConfigurationProviderIdentity struct {
	// The external ID used to estalish trust relationship with the AWS identity.
	ExternalId string `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	// The AWS identity principal.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`
}
