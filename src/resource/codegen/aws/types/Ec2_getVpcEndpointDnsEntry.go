package types

type Ec2_getVpcEndpointDnsEntry struct {
	// DNS name.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	// ID of the private hosted zone.
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`
}
