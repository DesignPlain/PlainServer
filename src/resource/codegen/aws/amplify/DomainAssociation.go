package amplify

import types "libds/aws/types"

type DomainAssociation struct {
	// Unique ID for an Amplify app.
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	// The type of SSL/TLS certificate to use for your custom domain. If you don't specify a certificate type, Amplify uses the default certificate that it provisions and manages for you.
	CertificateSettings types.Amplify_DomainAssociationCertificateSettings `json:"certificateSettings,omitempty" yaml:"certificateSettings,omitempty"`

	// Domain name for the domain association.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Enables the automated creation of subdomains for branches.
	EnableAutoSubDomain bool `json:"enableAutoSubDomain,omitempty" yaml:"enableAutoSubDomain,omitempty"`

	// Setting for the subdomain. Documented below.
	SubDomains []types.Amplify_DomainAssociationSubDomain `json:"subDomains,omitempty" yaml:"subDomains,omitempty"`

	// If enabled, the resource will wait for the domain association status to change to `PENDING_DEPLOYMENT` or `AVAILABLE`. Setting this to `false` will skip the process. Default: `true`.
	WaitForVerification bool `json:"waitForVerification,omitempty" yaml:"waitForVerification,omitempty"`
}
