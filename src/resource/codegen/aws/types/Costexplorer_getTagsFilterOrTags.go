package types

type Costexplorer_getTagsFilterOrTags struct {
	//
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	//
	MatchOptions []string `json:"matchOptions,omitempty" yaml:"matchOptions,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
