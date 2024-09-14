package types

type Imagebuilder_ImageRecipeComponentParameter struct {
	// The name of the component parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value for the named component parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
