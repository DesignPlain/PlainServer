package ec2

import types "libds/aws/types"

type TrafficMirrorFilterRule struct {
	// Destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty" yaml:"destinationCidrBlock,omitempty"`

	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange types.Ec2_TrafficMirrorFilterRuleSourcePortRange `json:"sourcePortRange,omitempty" yaml:"sourcePortRange,omitempty"`

	// Source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock string `json:"sourceCidrBlock,omitempty" yaml:"sourceCidrBlock,omitempty"`

	/*
	   Direction of traffic to be captured. Valid values are `ingress` and `egress`

	   Traffic mirror port range support following attributes:
	*/
	TrafficDirection string `json:"trafficDirection,omitempty" yaml:"trafficDirection,omitempty"`

	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId string `json:"trafficMirrorFilterId,omitempty" yaml:"trafficMirrorFilterId,omitempty"`

	// Description of the traffic mirror filter rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange types.Ec2_TrafficMirrorFilterRuleDestinationPortRange `json:"destinationPortRange,omitempty" yaml:"destinationPortRange,omitempty"`

	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol int `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction string `json:"ruleAction,omitempty" yaml:"ruleAction,omitempty"`

	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber int `json:"ruleNumber,omitempty" yaml:"ruleNumber,omitempty"`
}
