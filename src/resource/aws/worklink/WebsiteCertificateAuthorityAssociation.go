package worklink

type WebsiteCertificateAuthorityAssociation struct {
	// The certificate name to display.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The ARN of the fleet.
	FleetArn string `json:"fleetArn,omitempty" yaml:"fleetArn,omitempty"`

	// The root certificate of the Certificate Authority.
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`
}
