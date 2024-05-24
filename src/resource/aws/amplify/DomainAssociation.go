package amplify

import types "DesignSphere_Server/src/resource/aws/types"

type DomainAssociation struct {
	// Unique ID for an Amplify app.
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	// Domain name for the domain association.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Enables the automated creation of subdomains for branches.
	EnableAutoSubDomain bool `json:"enableAutoSubDomain,omitempty" yaml:"enableAutoSubDomain,omitempty"`

	// Setting for the subdomain. Documented below.
	SubDomains []types.Amplify_DomainAssociationSubDomain `json:"subDomains,omitempty" yaml:"subDomains,omitempty"`

	// If enabled, the resource will wait for the domain association status to change to `PENDING_DEPLOYMENT` or `AVAILABLE`. Setting this to `false` will skip the process. Default: `true`.
	WaitForVerification bool `json:"waitForVerification,omitempty" yaml:"waitForVerification,omitempty"`
}
