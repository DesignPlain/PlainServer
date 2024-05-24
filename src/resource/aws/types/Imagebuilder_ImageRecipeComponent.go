package types

type Imagebuilder_ImageRecipeComponent struct {
	// Amazon Resource Name (ARN) of the Image Builder Component to associate.
	ComponentArn string `json:"componentArn,omitempty" yaml:"componentArn,omitempty"`

	// Configuration block(s) for parameters to configure the component. Detailed below.
	Parameters []Imagebuilder_ImageRecipeComponentParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
