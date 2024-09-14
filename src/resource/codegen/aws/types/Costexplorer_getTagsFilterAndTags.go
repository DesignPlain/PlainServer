package types

type Costexplorer_getTagsFilterAndTags struct {
	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	//
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	//
	MatchOptions []string `json:"matchOptions,omitempty" yaml:"matchOptions,omitempty"`
}
