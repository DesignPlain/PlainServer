package route53

type ResolverDnsSecConfig struct {
	// The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`
}
