package types

type Certificateauthority_CaPoolIssuancePolicyBaselineValuesNameConstraints struct {
	// Indicates whether or not the name constraints are marked critical.
	Critical bool `json:"critical,omitempty" yaml:"critical,omitempty"`

	/*
	   Contains excluded DNS names. Any DNS name that can be
	   constructed by simply adding zero or more labels to
	   the left-hand side of the name satisfies the name constraint.
	   For example, `example.com`, `www.example.com`, `www.sub.example.com`
	   would satisfy `example.com` while `example1.com` does not.
	*/
	ExcludedDnsNames []string `json:"excludedDnsNames,omitempty" yaml:"excludedDnsNames,omitempty"`

	/*
	   Contains the excluded IP ranges. For IPv4 addresses, the ranges
	   are expressed using CIDR notation as specified in RFC 4632.
	   For IPv6 addresses, the ranges are expressed in similar encoding as IPv4
	   addresses.
	*/
	ExcludedIpRanges []string `json:"excludedIpRanges,omitempty" yaml:"excludedIpRanges,omitempty"`

	/*
	   Contains the excluded URIs that apply to the host part of the name.
	   The value can be a hostname or a domain with a
	   leading period (like `.example.com`)
	*/
	ExcludedUris []string `json:"excludedUris,omitempty" yaml:"excludedUris,omitempty"`

	/*
	   Contains permitted DNS names. Any DNS name that can be
	   constructed by simply adding zero or more labels to
	   the left-hand side of the name satisfies the name constraint.
	   For example, `example.com`, `www.example.com`, `www.sub.example.com`
	   would satisfy `example.com` while `example1.com` does not.
	*/
	PermittedDnsNames []string `json:"permittedDnsNames,omitempty" yaml:"permittedDnsNames,omitempty"`

	/*
	   Contains the permitted email addresses. The value can be a particular
	   email address, a hostname to indicate all email addresses on that host or
	   a domain with a leading period (e.g. `.example.com`) to indicate
	   all email addresses in that domain.
	*/
	PermittedEmailAddresses []string `json:"permittedEmailAddresses,omitempty" yaml:"permittedEmailAddresses,omitempty"`

	/*
	   Contains the permitted URIs that apply to the host part of the name.
	   The value can be a hostname or a domain with a
	   leading period (like `.example.com`)
	*/
	PermittedUris []string `json:"permittedUris,omitempty" yaml:"permittedUris,omitempty"`

	/*
	   Contains the excluded email addresses. The value can be a particular
	   email address, a hostname to indicate all email addresses on that host or
	   a domain with a leading period (e.g. `.example.com`) to indicate
	   all email addresses in that domain.
	*/
	ExcludedEmailAddresses []string `json:"excludedEmailAddresses,omitempty" yaml:"excludedEmailAddresses,omitempty"`

	/*
	   Contains the permitted IP ranges. For IPv4 addresses, the ranges
	   are expressed using CIDR notation as specified in RFC 4632.
	   For IPv6 addresses, the ranges are expressed in similar encoding as IPv4
	   addresses.
	*/
	PermittedIpRanges []string `json:"permittedIpRanges,omitempty" yaml:"permittedIpRanges,omitempty"`
}
