package types

type Ec2_getAmiFilter struct {
	// Name of the AMI that was provided during image creation.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
