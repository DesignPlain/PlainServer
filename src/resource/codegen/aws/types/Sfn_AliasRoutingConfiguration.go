package types

type Sfn_AliasRoutingConfiguration struct {
	// The Amazon Resource Name (ARN) of the state machine version.
	StateMachineVersionArn string `json:"stateMachineVersionArn,omitempty" yaml:"stateMachineVersionArn,omitempty"`

	// Percentage of traffic routed to the state machine version.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
