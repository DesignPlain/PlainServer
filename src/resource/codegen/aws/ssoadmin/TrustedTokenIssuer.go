package ssoadmin

import types "libds/aws/types"

type TrustedTokenIssuer struct {
	// Name of the trusted token issuer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A block that specifies settings that apply to the trusted token issuer, these change depending on the type you specify in `trusted_token_issuer_type`. Documented below.
	TrustedTokenIssuerConfiguration types.Ssoadmin_TrustedTokenIssuerTrustedTokenIssuerConfiguration `json:"trustedTokenIssuerConfiguration,omitempty" yaml:"trustedTokenIssuerConfiguration,omitempty"`

	/*
	   Specifies the type of the trusted token issuer. Valid values are `OIDC_JWT`

	   The following arguments are optional:
	*/
	TrustedTokenIssuerType string `json:"trustedTokenIssuerType,omitempty" yaml:"trustedTokenIssuerType,omitempty"`

	// A unique, case-sensitive ID that you provide to ensure the idempotency of the request. AWS generates a random value when not provided.
	ClientToken string `json:"clientToken,omitempty" yaml:"clientToken,omitempty"`

	// ARN of the instance of IAM Identity Center.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`
}
