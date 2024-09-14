package types

type Imagebuilder_getContainerRecipeComponent struct {
	// Set of parameters that are used to configure the component.
	Parameters []Imagebuilder_getContainerRecipeComponentParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// ARN of the Image Builder Component.
	ComponentArn string `json:"componentArn,omitempty" yaml:"componentArn,omitempty"`
}
