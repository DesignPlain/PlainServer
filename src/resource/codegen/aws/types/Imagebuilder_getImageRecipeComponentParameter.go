package types

type Imagebuilder_getImageRecipeComponentParameter struct {
	// Name of the image recipe.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the component parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
