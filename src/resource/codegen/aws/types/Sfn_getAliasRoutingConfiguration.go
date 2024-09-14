package types

type Sfn_getAliasRoutingConfiguration struct {
	//
	StateMachineVersionArn string `json:"stateMachineVersionArn,omitempty" yaml:"stateMachineVersionArn,omitempty"`

	//
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
