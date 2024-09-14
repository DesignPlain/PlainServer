package types

type Clouddomains_RegistrationDnsSettings struct {
	/*
	   Configuration for an arbitrary DNS provider.
	   Structure is documented below.
	*/
	CustomDns Clouddomains_RegistrationDnsSettingsCustomDns `json:"customDns,omitempty" yaml:"customDns,omitempty"`

	/*
	   The list of glue records for this Registration. Commonly empty.
	   Structure is documented below.
	*/
	GlueRecords []Clouddomains_RegistrationDnsSettingsGlueRecord `json:"glueRecords,omitempty" yaml:"glueRecords,omitempty"`
}
