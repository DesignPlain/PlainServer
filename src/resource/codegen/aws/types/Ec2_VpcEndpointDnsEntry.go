package types

type Ec2_VpcEndpointDnsEntry struct {
	// The DNS name.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	// The ID of the private hosted zone.
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`
}
