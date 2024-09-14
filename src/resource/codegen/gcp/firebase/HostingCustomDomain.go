package firebase

type HostingCustomDomain struct {
	/*
	   If true, Terraform will wait for DNS records to be fully resolved on the 'CustomDomain'. If false, Terraform will not
	   wait for DNS records on the 'CustomDomain'. Any issues in the 'CustomDomain' will be returned and stored in the
	   Terraform state.
	*/
	WaitDnsVerification bool `json:"waitDnsVerification,omitempty" yaml:"waitDnsVerification,omitempty"`

	/*
	   A field that lets you specify which SSL certificate type Hosting creates
	   for your domain name. Spark plan `CustomDomain`s only have access to the
	   `GROUPED` cert type, while Blaze plan can select any option.
	   Possible values are: `GROUPED`, `PROJECT_GROUPED`, `DEDICATED`.
	*/
	CertPreference string `json:"certPreference,omitempty" yaml:"certPreference,omitempty"`

	/*
	   The ID of the `CustomDomain`, which is the domain name you'd like to use with Firebase Hosting.


	   - - -
	*/
	CustomDomain string `json:"customDomain,omitempty" yaml:"customDomain,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A domain name that this CustomDomain should direct traffic towards. If
	   specified, Hosting will respond to requests against this CustomDomain
	   with an HTTP 301 code, and route traffic to the specified `redirect_target`
	   instead.
	*/
	RedirectTarget string `json:"redirectTarget,omitempty" yaml:"redirectTarget,omitempty"`

	// The ID of the site in which to create this custom domain association.
	SiteId string `json:"siteId,omitempty" yaml:"siteId,omitempty"`
}
