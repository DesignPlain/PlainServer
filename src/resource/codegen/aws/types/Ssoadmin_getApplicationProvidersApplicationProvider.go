package types

type Ssoadmin_getApplicationProvidersApplicationProvider struct {
	// Protocol that the application provider uses to perform federation. Valid values are `SAML` and `OAUTH`.
	FederationProtocol string `json:"federationProtocol,omitempty" yaml:"federationProtocol,omitempty"`

	// ARN of the application provider.
	ApplicationProviderArn string `json:"applicationProviderArn,omitempty" yaml:"applicationProviderArn,omitempty"`

	// An object describing how IAM Identity Center represents the application provider in the portal. See `display_data` below.
	DisplayDatas []Ssoadmin_getApplicationProvidersApplicationProviderDisplayData `json:"displayDatas,omitempty" yaml:"displayDatas,omitempty"`
}
