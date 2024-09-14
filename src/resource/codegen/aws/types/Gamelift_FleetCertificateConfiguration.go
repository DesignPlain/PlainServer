package types

type Gamelift_FleetCertificateConfiguration struct {
	// Indicates whether a TLS/SSL certificate is generated for a fleet. Valid values are `DISABLED` and `GENERATED`. Default value is `DISABLED`.
	CertificateType string `json:"certificateType,omitempty" yaml:"certificateType,omitempty"`
}
