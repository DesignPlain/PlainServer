package dax

type SubnetGroup struct {
	// A description of the subnet group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the subnet group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of VPC subnet IDs for the subnet group.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
