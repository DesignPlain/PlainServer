package types

type Ec2clientvpn_getEndpointAuthenticationOption struct {
	//
	ActiveDirectoryId string `json:"activeDirectoryId,omitempty" yaml:"activeDirectoryId,omitempty"`

	//
	RootCertificateChainArn string `json:"rootCertificateChainArn,omitempty" yaml:"rootCertificateChainArn,omitempty"`

	//
	SamlProviderArn string `json:"samlProviderArn,omitempty" yaml:"samlProviderArn,omitempty"`

	//
	SelfServiceSamlProviderArn string `json:"selfServiceSamlProviderArn,omitempty" yaml:"selfServiceSamlProviderArn,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
