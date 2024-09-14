package types

type Kendra_IndexDocumentMetadataConfigurationUpdate struct {
	// The name of the index field. Minimum length of 1. Maximum length of 30.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A block that provides manual tuning parameters to determine how the field affects the search results. Detailed below
	Relevance Kendra_IndexDocumentMetadataConfigurationUpdateRelevance `json:"relevance,omitempty" yaml:"relevance,omitempty"`

	// A block that provides information about how the field is used during a search. Documented below. Detailed below
	Search Kendra_IndexDocumentMetadataConfigurationUpdateSearch `json:"search,omitempty" yaml:"search,omitempty"`

	// The data type of the index field. Valid values are `STRING_VALUE`, `STRING_LIST_VALUE`, `LONG_VALUE`, `DATE_VALUE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
