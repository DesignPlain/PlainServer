package networkfirewall

import types "libds/aws/types"

type TlsInspectionConfiguration struct {
	//
	Timeouts types.Networkfirewall_TlsInspectionConfigurationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	/*
	   TLS inspection configuration block. Detailed below.

	   The following arguments are optional:
	*/
	TlsInspectionConfiguration types.Networkfirewall_TlsInspectionConfigurationTlsInspectionConfiguration `json:"tlsInspectionConfiguration,omitempty" yaml:"tlsInspectionConfiguration,omitempty"`

	// Description of the TLS inspection configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Encryption configuration block. Detailed below.
	EncryptionConfigurations []types.Networkfirewall_TlsInspectionConfigurationEncryptionConfiguration `json:"encryptionConfigurations,omitempty" yaml:"encryptionConfigurations,omitempty"`

	// Descriptive name of the TLS inspection configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
