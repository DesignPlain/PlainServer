package types

type Ec2clientvpn_EndpointAuthenticationOption struct {
	// The ID of the Active Directory to be used for authentication if type is `directory-service-authentication`.
	ActiveDirectoryId string `json:"activeDirectoryId,omitempty" yaml:"activeDirectoryId,omitempty"`

	// The ARN of the client certificate. The certificate must be signed by a certificate authority (CA) and it must be provisioned in AWS Certificate Manager (ACM). Only necessary when type is set to `certificate-authentication`.
	RootCertificateChainArn string `json:"rootCertificateChainArn,omitempty" yaml:"rootCertificateChainArn,omitempty"`

	// The ARN of the IAM SAML identity provider if type is `federated-authentication`.
	SamlProviderArn string `json:"samlProviderArn,omitempty" yaml:"samlProviderArn,omitempty"`

	// The ARN of the IAM SAML identity provider for the self service portal if type is `federated-authentication`.
	SelfServiceSamlProviderArn string `json:"selfServiceSamlProviderArn,omitempty" yaml:"selfServiceSamlProviderArn,omitempty"`

	// The type of client authentication to be used. Specify `certificate-authentication` to use certificate-based authentication, `directory-service-authentication` to use Active Directory authentication, or `federated-authentication` to use Federated Authentication via SAML 2.0.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
