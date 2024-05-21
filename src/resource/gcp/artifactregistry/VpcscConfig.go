package artifactregistry

type VpcscConfig struct {
	// The name of the location this config is located in.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The VPC SC policy for project and location.
	   Possible values are: `DENY`, `ALLOW`.
	*/
	VpcscPolicy string `json:"vpcscPolicy,omitempty" yaml:"vpcscPolicy,omitempty"`
}
