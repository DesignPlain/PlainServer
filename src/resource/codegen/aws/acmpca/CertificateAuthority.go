package acmpca

import types "libds/aws/types"

type CertificateAuthority struct {
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration types.Acmpca_CertificateAuthorityRevocationConfiguration `json:"revocationConfiguration,omitempty" yaml:"revocationConfiguration,omitempty"`

	// Key-value map of user-defined tags that are attached to the certificate authority. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to `GENERAL_PURPOSE`. Valid values: `GENERAL_PURPOSE` and `SHORT_LIVED_CERTIFICATE`.
	UsageMode string `json:"usageMode,omitempty" yaml:"usageMode,omitempty"`

	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration types.Acmpca_CertificateAuthorityCertificateAuthorityConfiguration `json:"certificateAuthorityConfiguration,omitempty" yaml:"certificateAuthorityConfiguration,omitempty"`

	// Whether the certificate authority is enabled or disabled. Defaults to `true`. Can only be disabled if the CA is in an `ACTIVE` state.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Cryptographic key management compliance standard used for handling CA keys. Defaults to `FIPS_140_2_LEVEL_3_OR_HIGHER`. Valid values: `FIPS_140_2_LEVEL_3_OR_HIGHER` and `FIPS_140_2_LEVEL_2_OR_HIGHER`. Supported standard for each region can be found in the [Storage and security compliance of AWS Private CA private keys Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
	KeyStorageSecurityStandard string `json:"keyStorageSecurityStandard,omitempty" yaml:"keyStorageSecurityStandard,omitempty"`

	// Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays int `json:"permanentDeletionTimeInDays,omitempty" yaml:"permanentDeletionTimeInDays,omitempty"`
}
