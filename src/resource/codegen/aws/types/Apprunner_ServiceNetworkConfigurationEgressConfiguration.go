package types

type Apprunner_ServiceNetworkConfigurationEgressConfiguration struct {
	// The type of egress configuration. Valid values are: `DEFAULT` and `VPC`.
	EgressType string `json:"egressType,omitempty" yaml:"egressType,omitempty"`

	// The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to associate with your App Runner service. Only valid when `EgressType = VPC`.
	VpcConnectorArn string `json:"vpcConnectorArn,omitempty" yaml:"vpcConnectorArn,omitempty"`
}
