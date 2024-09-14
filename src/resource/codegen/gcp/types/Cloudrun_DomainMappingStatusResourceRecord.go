package types

type Cloudrun_DomainMappingStatusResourceRecord struct {
	// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   (Output)
	   Data for this record. Values vary by record type, as defined in RFC 1035
	   (section 5) and RFC 1034 (section 3.6.1).
	*/
	Rrdata string `json:"rrdata,omitempty" yaml:"rrdata,omitempty"`

	/*
	   Resource record type. Example: `AAAA`.
	   Possible values are: `A`, `AAAA`, `CNAME`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
