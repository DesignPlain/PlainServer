package types

type Wafregional_RulePredicate struct {
	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	//
	DataId string `json:"dataId,omitempty" yaml:"dataId,omitempty"`

	//
	Negated bool `json:"negated,omitempty" yaml:"negated,omitempty"`
}
