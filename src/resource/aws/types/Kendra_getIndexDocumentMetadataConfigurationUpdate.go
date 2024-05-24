package types

type Kendra_getIndexDocumentMetadataConfigurationUpdate struct {
	// Name of the index field. Minimum length of 1. Maximum length of 30.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Block that provides manual tuning parameters to determine how the field affects the search results. Documented below.
	Relevances []Kendra_getIndexDocumentMetadataConfigurationUpdateRelevance `json:"relevances,omitempty" yaml:"relevances,omitempty"`

	// Block that provides information about how the field is used during a search. Documented below.
	Searches []Kendra_getIndexDocumentMetadataConfigurationUpdateSearch `json:"searches,omitempty" yaml:"searches,omitempty"`

	// Data type of the index field. Valid values are `STRING_VALUE`, `STRING_LIST_VALUE`, `LONG_VALUE`, `DATE_VALUE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
