package types

type Verifiedaccess_InstanceVerifiedAccessTrustProvider struct {
	// A description for the AWS Verified Access Instance.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The type of device-based trust provider.
	DeviceTrustProviderType string `json:"deviceTrustProviderType,omitempty" yaml:"deviceTrustProviderType,omitempty"`

	// The type of trust provider (user- or device-based).
	TrustProviderType string `json:"trustProviderType,omitempty" yaml:"trustProviderType,omitempty"`

	// The type of user-based trust provider.
	UserTrustProviderType string `json:"userTrustProviderType,omitempty" yaml:"userTrustProviderType,omitempty"`

	// The ID of the trust provider.
	VerifiedAccessTrustProviderId string `json:"verifiedAccessTrustProviderId,omitempty" yaml:"verifiedAccessTrustProviderId,omitempty"`
}
