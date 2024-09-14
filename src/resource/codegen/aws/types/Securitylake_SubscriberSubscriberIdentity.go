package types

type Securitylake_SubscriberSubscriberIdentity struct {
	// The AWS Regions where Security Lake is automatically enabled.
	ExternalId string `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	// Provides encryption details of Amazon Security Lake object.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`
}
