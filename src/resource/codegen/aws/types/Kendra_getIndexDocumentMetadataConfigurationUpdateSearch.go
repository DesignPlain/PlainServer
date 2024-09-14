package types

type Kendra_getIndexDocumentMetadataConfigurationUpdateSearch struct {
	// Determines whether the field is used in the search. If the Searchable field is true, you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is `true` for `string` fields and `false` for `number` and `date` fields.
	Searchable bool `json:"searchable,omitempty" yaml:"searchable,omitempty"`

	// Determines whether the field can be used to sort the results of a query. If you specify sorting on a field that does not have Sortable set to true, Amazon Kendra returns an exception. The default is `false`.
	Sortable bool `json:"sortable,omitempty" yaml:"sortable,omitempty"`

	// Determines whether the field is returned in the query response. The default is `true`.
	Displayable bool `json:"displayable,omitempty" yaml:"displayable,omitempty"`

	// Whether the field can be used to create search facets, a count of results for each value in the field. The default is `false`.
	Facetable bool `json:"facetable,omitempty" yaml:"facetable,omitempty"`
}
