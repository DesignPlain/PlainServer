package apigee

import types "DesignSphere_Server/src/resource/gcp/types"

type KeystoresAliasesKeyCertFile struct {
	/*
	   Chain of certificates under this alias.
	   Structure is documented below.
	*/
	CertsInfo types.Apigee_KeystoresAliasesKeyCertFileCertsInfo `json:"certsInfo,omitempty" yaml:"certsInfo,omitempty"`

	// Environment associated with the alias
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`

	// Private Key content, omit if uploading to truststore
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Keystore Name
	Keystore string `json:"keystore,omitempty" yaml:"keystore,omitempty"`

	// Organization ID associated with the alias, without organization/ prefix
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// Password for the Private Key if it's encrypted
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Alias Name
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	/*
	   Cert content


	   - - -
	*/
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`
}
