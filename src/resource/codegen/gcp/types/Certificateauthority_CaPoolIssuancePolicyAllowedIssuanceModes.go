package types

type Certificateauthority_CaPoolIssuancePolicyAllowedIssuanceModes struct {
	// When true, allows callers to create Certificates by specifying a CertificateConfig.
	AllowConfigBasedIssuance bool `json:"allowConfigBasedIssuance,omitempty" yaml:"allowConfigBasedIssuance,omitempty"`

	// When true, allows callers to create Certificates by specifying a CSR.
	AllowCsrBasedIssuance bool `json:"allowCsrBasedIssuance,omitempty" yaml:"allowCsrBasedIssuance,omitempty"`
}
