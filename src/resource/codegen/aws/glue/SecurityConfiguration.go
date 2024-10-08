package glue

import types "libds/aws/types"

type SecurityConfiguration struct {
	// Configuration block containing encryption configuration. Detailed below.
	EncryptionConfiguration types.Glue_SecurityConfigurationEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// Name of the security configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
