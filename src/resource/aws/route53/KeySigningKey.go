package route53

type KeySigningKey struct {
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
	KeyManagementServiceArn string `json:"keyManagementServiceArn,omitempty" yaml:"keyManagementServiceArn,omitempty"`

	/*
	   Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
