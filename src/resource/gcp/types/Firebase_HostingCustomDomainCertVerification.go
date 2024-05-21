package types

type Firebase_HostingCustomDomainCertVerification struct {
	/*
	   A `TXT` record to add to your DNS records that confirms your intent to
	   let Hosting create an SSL cert for your domain name.
	   Structure is documented below.
	*/
	Dns Firebase_HostingCustomDomainCertVerificationDns `json:"dns,omitempty" yaml:"dns,omitempty"`

	/*
	   A file to add to your existing, non-Hosting hosting service that confirms
	   your intent to let Hosting create an SSL cert for your domain name.
	   Structure is documented below.
	*/
	Http Firebase_HostingCustomDomainCertVerificationHttp `json:"http,omitempty" yaml:"http,omitempty"`
}
