package types

type Imagebuilder_getContainerRecipesFilter struct {
	// Name of the filter field. Valid values can be found in the [Image Builder ListContainerRecipes API Reference](https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_ListContainerRecipes.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
