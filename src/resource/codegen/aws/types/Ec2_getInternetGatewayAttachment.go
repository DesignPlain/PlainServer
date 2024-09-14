package types

type Ec2_getInternetGatewayAttachment struct {
	// Current state of the attachment between the gateway and the VPC. Present only if a VPC is attached
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// ID of an attached VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
