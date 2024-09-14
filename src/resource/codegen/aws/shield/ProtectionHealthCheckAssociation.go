package shield

type ProtectionHealthCheckAssociation struct {
	// The ARN (Amazon Resource Name) of the Route53 Health Check resource which will be associated to the protected resource.
	HealthCheckArn string `json:"healthCheckArn,omitempty" yaml:"healthCheckArn,omitempty"`

	// The ID of the protected resource.
	ShieldProtectionId string `json:"shieldProtectionId,omitempty" yaml:"shieldProtectionId,omitempty"`
}
