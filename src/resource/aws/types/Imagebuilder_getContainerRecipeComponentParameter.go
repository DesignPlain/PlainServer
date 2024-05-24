package types

type Imagebuilder_getContainerRecipeComponentParameter struct {
	// Name of the container recipe.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the component parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
