package apigee

type KeystoresAliasesPkcs12 struct {
	// Environment associated with the alias
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`

	/*
	   PKCS12 file content

	   - - -
	*/
	File string `json:"file,omitempty" yaml:"file,omitempty"`

	// Hash of the pkcs file
	Filehash string `json:"filehash,omitempty" yaml:"filehash,omitempty"`

	// Keystore Name
	Keystore string `json:"keystore,omitempty" yaml:"keystore,omitempty"`

	// Organization ID associated with the alias, without organization/ prefix
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// Password for the PKCS12 file if it's encrypted
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Alias Name
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`
}
