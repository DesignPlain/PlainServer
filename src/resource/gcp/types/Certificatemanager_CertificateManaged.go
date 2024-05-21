package types

type Certificatemanager_CertificateManaged struct {
	// Authorizations that will be used for performing domain authorization. Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	DnsAuthorizations []string `json:"dnsAuthorizations,omitempty" yaml:"dnsAuthorizations,omitempty"`

	/*
	   The domains for which a managed SSL certificate will be generated.
	   Wildcard domains are only supported with DNS challenge resolution
	*/
	Domains []string `json:"domains,omitempty" yaml:"domains,omitempty"`

	/*
	   The resource name for a CertificateIssuanceConfig used to configure private PKI certificates in the format projects/-/locations/-/certificateIssuanceConfigs/-.
	   If this field is not set, the certificates will instead be publicly signed as documented at https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.
	   Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	*/
	IssuanceConfig string `json:"issuanceConfig,omitempty" yaml:"issuanceConfig,omitempty"`

	/*
	   (Output)
	   Information about issues with provisioning this Managed Certificate.
	   Structure is documented below.
	*/
	ProvisioningIssues []Certificatemanager_CertificateManagedProvisioningIssue `json:"provisioningIssues,omitempty" yaml:"provisioningIssues,omitempty"`

	/*
	   (Output)
	   State of the domain for managed certificate issuance.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   (Output)
	   Detailed state of the latest authorization attempt for each domain
	   specified for this Managed Certificate.
	   Structure is documented below.


	   <a name="nested_provisioning_issue"></a>The `provisioning_issue` block contains:
	*/
	AuthorizationAttemptInfos []Certificatemanager_CertificateManagedAuthorizationAttemptInfo `json:"authorizationAttemptInfos,omitempty" yaml:"authorizationAttemptInfos,omitempty"`
}
