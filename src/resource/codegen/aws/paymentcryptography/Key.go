package paymentcryptography

import types "libds/aws/types"

type Key struct {
	/*
	   Role of the key, the algorithm it supports, and the cryptographic operations allowed with the key.

	   The following arguments are optional:
	*/
	KeyAttributes types.Paymentcryptography_KeyKeyAttributes `json:"keyAttributes,omitempty" yaml:"keyAttributes,omitempty"`

	// Algorithm that AWS Payment Cryptography uses to calculate the key check value (KCV).
	KeyCheckValueAlgorithm string `json:"keyCheckValueAlgorithm,omitempty" yaml:"keyCheckValueAlgorithm,omitempty"`

	// Map of tags assigned to the WorkSpaces Connection Alias. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Paymentcryptography_KeyTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	//
	DeletionWindowInDays int `json:"deletionWindowInDays,omitempty" yaml:"deletionWindowInDays,omitempty"`

	// Whether to enable the key.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Whether the key is exportable from the service.
	Exportable bool `json:"exportable,omitempty" yaml:"exportable,omitempty"`
}
