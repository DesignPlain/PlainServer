package types

type Sfn_getAliasRoutingConfiguration struct {
	//
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`

	//
	StateMachineVersionArn string `json:"stateMachineVersionArn,omitempty" yaml:"stateMachineVersionArn,omitempty"`
}
