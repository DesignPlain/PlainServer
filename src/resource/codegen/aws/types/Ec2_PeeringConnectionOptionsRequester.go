package types

type Ec2_PeeringConnectionOptionsRequester struct {
	// Allow a local VPC to resolve public DNS hostnames to private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVpcDnsResolution bool `json:"allowRemoteVpcDnsResolution,omitempty" yaml:"allowRemoteVpcDnsResolution,omitempty"`
}
