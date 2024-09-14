package ec2

type InternetGatewayAttachment struct {
	// The ID of the internet gateway.
	InternetGatewayId string `json:"internetGatewayId,omitempty" yaml:"internetGatewayId,omitempty"`

	// The ID of the VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
