package types

type Imagebuilder_ImageOutputResource struct {
	// Set of objects with each Amazon Machine Image (AMI) created.
	Amis []Imagebuilder_ImageOutputResourceAmi `json:"amis,omitempty" yaml:"amis,omitempty"`

	// Set of objects with each container image created and stored in the output repository.
	Containers []Imagebuilder_ImageOutputResourceContainer `json:"containers,omitempty" yaml:"containers,omitempty"`
}
