package types

type Certificateauthority_CertificateCertificateDescriptionX509DescriptionCaOption struct {
	// When true, the "CA" in Basic Constraints extension will be set to true.
	IsCa bool `json:"isCa,omitempty" yaml:"isCa,omitempty"`

	/*
	   Refers to the "path length constraint" in Basic Constraints extension. For a CA certificate, this value describes the depth of
	   subordinate CA certificates that are allowed. If this value is less than 0, the request will fail.
	*/
	MaxIssuerPathLength int `json:"maxIssuerPathLength,omitempty" yaml:"maxIssuerPathLength,omitempty"`
}
