package types

type Imagebuilder_ImageOutputResourceAmi struct {
	// Account identifier of the AMI.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Description of the AMI.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Identifier of the AMI.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Name of the AMI.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Region of the container image.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
