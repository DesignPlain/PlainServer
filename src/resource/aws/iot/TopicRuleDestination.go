package iot

import types "DesignSphere_Server/src/resource/aws/types"

type TopicRuleDestination struct {
	// Whether or not to enable the destination. Default: `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Configuration of the virtual private cloud (VPC) connection. For more info, see the [AWS documentation](https://docs.aws.amazon.com/iot/latest/developerguide/vpc-rule-action.html).
	VpcConfiguration types.Iot_TopicRuleDestinationVpcConfiguration `json:"vpcConfiguration,omitempty" yaml:"vpcConfiguration,omitempty"`
}
