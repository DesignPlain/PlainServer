package types

type Apprunner_VpcIngressConnectionIngressVpcConfiguration struct {
	// The ID of the VPC endpoint that your App Runner service connects to.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// The ID of the VPC that is used for the VPC endpoint.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
