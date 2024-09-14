package types

type Msk_getBrokerNodesNodeInfoList struct {
	// ARN of the node
	NodeArn string `json:"nodeArn,omitempty" yaml:"nodeArn,omitempty"`

	// Attached elastic network interface of the broker
	AttachedEniId string `json:"attachedEniId,omitempty" yaml:"attachedEniId,omitempty"`

	// ID of the broker
	BrokerId float64 `json:"brokerId,omitempty" yaml:"brokerId,omitempty"`

	// Client subnet to which this broker node belongs
	ClientSubnet string `json:"clientSubnet,omitempty" yaml:"clientSubnet,omitempty"`

	// The client virtual private cloud (VPC) IP address
	ClientVpcIpAddress string `json:"clientVpcIpAddress,omitempty" yaml:"clientVpcIpAddress,omitempty"`

	// Set of endpoints for accessing the broker. This does not include ports
	Endpoints []string `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
}
