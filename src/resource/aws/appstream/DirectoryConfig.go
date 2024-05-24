package appstream

import types "DesignSphere_Server/src/resource/aws/types"

type DirectoryConfig struct {
	// Distinguished names of the organizational units for computer accounts.
	OrganizationalUnitDistinguishedNames []string `json:"organizationalUnitDistinguishedNames,omitempty" yaml:"organizationalUnitDistinguishedNames,omitempty"`

	// Configuration block for the name of the directory and organizational unit (OU) to use to join the directory config to a Microsoft Active Directory domain. See `service_account_credentials` below.
	ServiceAccountCredentials types.Appstream_DirectoryConfigServiceAccountCredentials `json:"serviceAccountCredentials,omitempty" yaml:"serviceAccountCredentials,omitempty"`

	// Fully qualified name of the directory.
	DirectoryName string `json:"directoryName,omitempty" yaml:"directoryName,omitempty"`
}
