package verifiedaccess

import types "libds/aws/types"

type TrustProvider struct {
	// The identifier to be used when working with policy rules.
	PolicyReferenceName string `json:"policyReferenceName,omitempty" yaml:"policyReferenceName,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The type of trust provider can be either user or device-based.

	   The following arguments are optional:
	*/
	TrustProviderType string `json:"trustProviderType,omitempty" yaml:"trustProviderType,omitempty"`

	// The type of user-based trust provider.
	UserTrustProviderType string `json:"userTrustProviderType,omitempty" yaml:"userTrustProviderType,omitempty"`

	// A description for the AWS Verified Access trust provider.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A block of options for device identity based trust providers.
	DeviceOptions types.Verifiedaccess_TrustProviderDeviceOptions `json:"deviceOptions,omitempty" yaml:"deviceOptions,omitempty"`

	// The type of device-based trust provider.
	DeviceTrustProviderType string `json:"deviceTrustProviderType,omitempty" yaml:"deviceTrustProviderType,omitempty"`

	// The OpenID Connect details for an oidc-type, user-identity based trust provider.
	OidcOptions types.Verifiedaccess_TrustProviderOidcOptions `json:"oidcOptions,omitempty" yaml:"oidcOptions,omitempty"`
}
