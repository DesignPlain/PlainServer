package types

type Ec2_getAmiIdsFilter struct {
	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
