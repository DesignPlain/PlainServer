package types

type Acmpca_CertificateValidity struct {
	// Determines how `value` is interpreted. Valid values: `DAYS`, `MONTHS`, `YEARS`, `ABSOLUTE`, `END_DATE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// If `type` is `DAYS`, `MONTHS`, or `YEARS`, the relative time until the certificate expires. If `type` is `ABSOLUTE`, the date in seconds since the Unix epoch. If `type` is `END_DATE`, the  date in RFC 3339 format.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
