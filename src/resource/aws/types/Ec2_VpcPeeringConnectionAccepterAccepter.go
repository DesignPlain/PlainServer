package types

type Ec2_VpcPeeringConnectionAccepterAccepter struct {
	/*
	   Indicates whether a local VPC can resolve public DNS hostnames to
	   private IP addresses when queried from instances in a peer VPC.
	*/
	AllowRemoteVpcDnsResolution bool `json:"allowRemoteVpcDnsResolution,omitempty" yaml:"allowRemoteVpcDnsResolution,omitempty"`
}
