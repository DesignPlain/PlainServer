package route53

type HostedZoneDnsSec struct {
	/*
	   Identifier of the Route 53 Hosted Zone.

	   The following arguments are optional:
	*/
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`

	// Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
	SigningStatus string `json:"signingStatus,omitempty" yaml:"signingStatus,omitempty"`
}
