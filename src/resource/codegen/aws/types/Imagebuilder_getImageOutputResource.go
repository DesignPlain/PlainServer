package types

type Imagebuilder_getImageOutputResource struct {
	// Set of objects with each Amazon Machine Image (AMI) created.
	Amis []Imagebuilder_getImageOutputResourceAmi `json:"amis,omitempty" yaml:"amis,omitempty"`

	// Set of objects with each container image created and stored in the output repository.
	Containers []Imagebuilder_getImageOutputResourceContainer `json:"containers,omitempty" yaml:"containers,omitempty"`
}
