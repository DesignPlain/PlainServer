package types

type Acm_CertificateRenewalSummary struct {
	// The status of ACM's managed renewal of the certificate
	RenewalStatus string `json:"renewalStatus,omitempty" yaml:"renewalStatus,omitempty"`

	// The reason that a renewal request was unsuccessful or is pending
	RenewalStatusReason string `json:"renewalStatusReason,omitempty" yaml:"renewalStatusReason,omitempty"`

	//
	UpdatedAt string `json:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`
}
