package clouddomains

import types "DesignSphere_Server/src/resource/gcp/types"

type Registration struct {
	/*
	   Required. Yearly price to register or renew the domain. The value that should be put here can be obtained from
	   registrations.retrieveRegisterParameters or registrations.searchDomains calls.
	   Structure is documented below.
	*/
	YearlyPrice types.Clouddomains_RegistrationYearlyPrice `json:"yearlyPrice,omitempty" yaml:"yearlyPrice,omitempty"`

	// The list of contact notices that the caller acknowledges. Possible value is PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT
	ContactNotices []string `json:"contactNotices,omitempty" yaml:"contactNotices,omitempty"`

	/*
	   Required. Settings for contact information linked to the Registration.
	   Structure is documented below.
	*/
	ContactSettings types.Clouddomains_RegistrationContactSettings `json:"contactSettings,omitempty" yaml:"contactSettings,omitempty"`

	/*
	   Settings controlling the DNS configuration of the Registration.
	   Structure is documented below.
	*/
	DnsSettings types.Clouddomains_RegistrationDnsSettings `json:"dnsSettings,omitempty" yaml:"dnsSettings,omitempty"`

	/*
	   Set of labels associated with the Registration.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Required. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The list of domain notices that you acknowledge. Possible value is HSTS_PRELOADED
	DomainNotices []string `json:"domainNotices,omitempty" yaml:"domainNotices,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Settings for management of the Registration, including renewal, billing, and transfer
	   Structure is documented below.
	*/
	ManagementSettings types.Clouddomains_RegistrationManagementSettings `json:"managementSettings,omitempty" yaml:"managementSettings,omitempty"`
}
