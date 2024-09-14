package types

type Ec2_TrafficMirrorFilterRuleSourcePortRange struct {
	// Ending port of the range
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// Starting port of the range
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`
}
