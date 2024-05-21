package types

type Firebase_HostingCustomDomainCertVerificationDnsDiscoveredRecord struct {
	// The domain name the record pertains to, e.g. `foo.bar.com.`.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	/*
	   The data of the record. The meaning of the value depends on record type:
	   - A and AAAA: IP addresses for the domain name.
	   - CNAME: Another domain to check for records.
	   - TXT: Arbitrary text strings associated with the domain name. Hosting
	   uses TXT records to determine a which Firebase Projects have
	   permission to act on the domain name's behalf.
	   - CAA: The record's flags, tag, and value, e.g. `0 issue "pki.goog"`.
	*/
	Rdata string `json:"rdata,omitempty" yaml:"rdata,omitempty"`

	// Indicates the a required action for this record.
	RequiredAction string `json:"requiredAction,omitempty" yaml:"requiredAction,omitempty"`

	// The record's type, which determines what data the record contains.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
